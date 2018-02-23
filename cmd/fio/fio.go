package main

import (
	"fmt"
	"os"
	"path"

	fireblock "github.com/fireblock/go-fireblock"
	"github.com/fireblock/go-fireblock/common"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("fio", "fireblock.io CLI (verify and sign)")

	jverbose = app.Flag("json", "Output in JSON format").Short('j').Bool()
	server   = app.Flag("server", "Default to fireblock.io").Short('s').Default("fireblock.io").String()

	verifyCmd = app.Command("verify", "verify a file")
	cardID    = verifyCmd.Flag("card-id", "Set card id").Short('c').Default("0x0").String()
	userID    = verifyCmd.Flag("user-id", "Set user id").Short('u').Default("0x0").String()
	fverify   = verifyCmd.Arg("file", "File to verify.").Required().ExistingFile()

	// config file ~/.config/fireblock/fio.yaml overrided
	signCmd    = app.Command("sign", "sign a file")
	token      = signCmd.Flag("token", "Set token").Short('t').Default("").String()
	fkey       = signCmd.Flag("key", "Set token").Short('k').Required().ExistingFile()
	passphrase = signCmd.Flag("passphrase", "Set token").Short('p').Default("").String()
	fsign      = signCmd.Arg("file", "File to sign.").Required().ExistingFile()
)

func exit(msg string) {
	fmt.Printf("%s\n", msg)
	os.Exit(1)
}

func sign() {
	// check token
	if token == nil || *token == "" {
		exit("Missing token! add -t tokenValue")
	}
	// check fio file
	if fkey == nil || *fkey == "" {
		exit("Missing key! add -k filepath")
	}
	// load fio file
	keypath := *fkey
	keyuid, privkey, _, err := common.LoadFioFile(keypath)
	if err != nil {
		exit("Invalid fio file")
	}
	// compute sha256
	filepath := *fsign
	sha256, err := common.Sha256File(filepath)
	if err != nil {
		exit(fmt.Sprintf("Cannot compute sha256 on %s\n", filepath))
	}
	// get metadata
	metadata, err := common.MetadataFile(filepath)
	if err != nil {
		exit(fmt.Sprintf("Cannot read metadata on %s\n", filepath))
	}
	// get Card
	card, err := common.HTTPCard(keyuid, *token)
	if err != nil {
		exit("Invalid Token")
	}
	// create message
	message := sha256 + "||" + common.Keccak256(card)
	// create signature
	signature := ""
	ktype := common.B32Type(keyuid)
	if ktype == "pgp" {
		signature, err = common.PGPSign(message, privkey, *passphrase)
		if err != nil {
			exit("Can't sign")
		}
	} else if ktype == "ecdsa" {
		signature, err = common.ECDSASign(privkey, message)
		if err != nil {
			exit("Can't sign")
		}
	} else {
		exit(fmt.Sprintf("Invalid key format %s\n", ktype))
	}
	// sign
	_, err = common.HTTPSign(*token, sha256, keyuid, signature, metadata)
	if err != nil {
		exit("Can't sign")
	}
	filename := path.Base(filepath)
	fmt.Print(filename)
	fmt.Print(sha256)
	fmt.Print(metadata)
}

func verify() {
	if fverify == nil {
		exit("Missing file")
	}

	if *cardID == "0x0" && *userID == "0x0" {
		exit("Use -u user-id or -c card-id")
	}
	filepath := *fverify
	filename := path.Base(filepath)
	sha256, err := common.Sha256File(filepath)
	if err != nil {
		fmt.Printf("Cannot compute sha256 on %s\n", filepath)
		os.Exit(1)
	}
	if *userID != "0x0" {
		// verify by useruid
		GVerify(*server, filename, sha256, *userID, *jverbose)
	} else {
		// verify by cardId
		CVerify(*server, filename, sha256, *cardID, *jverbose)
	}
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(fireblock.Version).Author(fireblock.Author)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case signCmd.FullCommand():
		sign()
	case verifyCmd.FullCommand():
		verify()
	}
}
