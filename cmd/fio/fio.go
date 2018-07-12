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
	server   = app.Flag("server", "Default to fireblock.io").Short('s').Default("dev.fireblock.io").String()

	verifyCmd = app.Command("verify", "verify a file")
	projectID = verifyCmd.Flag("project-id", "Set project id").Short('p').Default("0x0").String()
	userID    = verifyCmd.Flag("user-id", "Set user id").Short('u').Default("0x0").String()
	fverify   = verifyCmd.Arg("file", "File to verify.").Required().ExistingFile()

	signCmd    = app.Command("sign", "sign a file")
	signKey    = signCmd.Flag("key", "private key in base64url").Short('k').Default("").String()
	signFio    = signCmd.Flag("fio", "path to fio file").Short('f').ExistingFile()
	passphrase = signCmd.Flag("passphrase", "passphrase (for PGP private key)").Short('p').Default("").String()
	fsign      = signCmd.Arg("file", "File to sign.").Required().ExistingFile()

	versionCmd = app.Command("version", "versio of fio (https://fireblock.io)")
)

func exit(msg string) {
	fmt.Printf("%s\n", msg)
	os.Exit(1)
}

func sign() {
	// Read the private key (signFio or signKey)
	var keyuid, privkey string
	var err error
	// check fio file
	if (signFio == nil || *signFio == "") && (*signKey == "") {
		exit("Missing private key! add --fio filepath or --key privatekey")
	} else if signFio != nil && len(*signFio) > 1 {
		keypath := *signFio
		// load fio file
		keyuid, privkey, _, err = common.LoadFioFile(keypath)
		if err != nil {
			exit("Invalid fio file")
		}
	} else if signKey != nil && len(*signKey) > 1 {
		keyuid, privkey, err = common.LoadB64U(*signKey)
	} else {
		exit("Missing private key! add --fio filepath or --key privatekey with correct values")
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
	// create message
	message := sha256 + "||" + keyuid
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
	_, err = createCertificate(*server, sha256, keyuid, signature, metadata)
	if err != nil {
		fmt.Println(err)
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

	if *projectID == "0x0" && *userID == "0x0" {
		exit("Use -u user-id or -p project-id")
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
		uVerify(*server, filename, sha256, *userID, *jverbose)
	} else {
		// verify by cardId
		pVerify(*server, filename, sha256, *projectID, *jverbose)
	}
}

func version() {
	fmt.Printf("fio %s ( info at https://fireblock.io )\n", fireblock.Version)
	os.Exit(1)
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(fireblock.Version).Author(fireblock.Author)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case signCmd.FullCommand():
		sign()
	case verifyCmd.FullCommand():
		verify()
	case versionCmd.FullCommand():
		version()
	default:
		exit("unknown command. Use --help")
	}
}
