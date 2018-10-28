// Copyright 2015-2017 Fireblock.
// This file is part of Fireblock.

// Fireblock is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Fireblock is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Fireblock.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"path"

	fireblock "github.com/fireblock/go-fireblock"
	"github.com/fireblock/go-fireblock/fireblocklib"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("fio", "fireblock.io CLI (verify and sign)")

	jverbose = app.Flag("json", "Output in JSON format").Short('j').Bool()
	server   = app.Flag("server", "Default to https://fireblock.io").Short('s').Default("https://fireblock.io").String() // dev.fireblock.io

	verifyCmd = app.Command("verify", "verify a file")
	projectID = verifyCmd.Flag("project-id", "Set project id").Short('p').Default("0x0").String()
	userID    = verifyCmd.Flag("user-id", "Set user id").Short('u').Default("").String()
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

func signFunction() {
	// Read the private key (signFio or signKey)
	var keyuid, privkey string
	var err error
	// check fio file
	if (signFio == nil || *signFio == "") && (*signKey == "") {
		exit("Missing private key! add --fio filepath or --key privatekey")
	} else if signFio != nil && len(*signFio) > 1 {
		keypath := *signFio
		// load fio file
		_, keyuid, privkey, _, err = fireblocklib.LoadFioFile(keypath)
		if err != nil {
			exit("Invalid fio file")
		}
	} else if signKey != nil && len(*signKey) > 1 {
		keyuid, privkey, err = fireblocklib.LoadB64U(*signKey)
	} else {
		exit("Missing private key! add --fio filepath or --key privatekey with correct values")
	}

	// compute sha256
	filepath := *fsign
	sha256, err := fireblocklib.Sha256File(filepath)
	if err != nil {
		exit(fmt.Sprintf("Cannot compute sha256 on %s\n", filepath))
	}
	// get metadata
	metadata, err := fireblocklib.MetadataFile(filepath)
	if err != nil {
		exit(fmt.Sprintf("Cannot read metadata on %s\n", filepath))
	}
	metadataSID := fireblocklib.Keccak256(metadata)
	// create message
	message := sha256 + "||" + keyuid
	messageSignature := metadataSID + "||" + sha256 + "||" + keyuid
	// create signature
	signature := ""
	metadataSignature := ""
	ktype := fireblocklib.B32Type(keyuid)
	if ktype == "pgp" {
		signature, err = fireblocklib.PGPSign(message, privkey, *passphrase)
		if err != nil {
			exit("Can't sign")
		}
		metadataSignature, err = fireblocklib.PGPSign(messageSignature, privkey, *passphrase)
		if err != nil {
			exit("Can't sign")
		}
	} else if ktype == "ecdsa" {
		signature, err = fireblocklib.ECDSASign(privkey, message)
		if err != nil {
			exit("Can't sign")
		}
		metadataSignature, err = fireblocklib.ECDSASign(privkey, messageSignature)
		if err != nil {
			exit("Can't sign")
		}
	} else {
		exit(fmt.Sprintf("Invalid key format %s\n", ktype))
	}
	// sign
	_, err = createCertificate(*server, sha256, ktype, keyuid, signature, metadata, metadataSignature)
	if err != nil {
		fmt.Println(err)
		exit("Can't sign")
	}
	filename := path.Base(filepath)
	fmt.Print(filename)
	fmt.Print(sha256)
	fmt.Print(metadata)
}

func verifyFunction() {
	if fverify == nil {
		exit("Missing file")
	}
	filepath := *fverify
	filename := path.Base(filepath)
	sha256, err := fireblocklib.Sha256File(filepath)
	if err != nil {
		fmt.Printf("Cannot compute sha256 on %s\n", filepath)
		os.Exit(1)
	}
	if *userID != "" {
		// verify by useruid
		userVerify(*server, filename, sha256, *userID, *jverbose)
	} else if *projectID != "0x0" {
		// verify by cardId
		projectVerify(*server, filename, sha256, *projectID, *jverbose)
	} else {
		// verify with the first key to register that hash
		verify(*server, filename, sha256, *jverbose)
	}
}

func version() {
	fmt.Printf("fio %s ( info at https://fireblock.io )\n", fireblock.Version)
	os.Exit(0)
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(fireblock.Version).Author(fireblock.Author)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case signCmd.FullCommand():
		signFunction()
	case verifyCmd.FullCommand():
		verifyFunction()
	case versionCmd.FullCommand():
		version()
	default:
		exit("unknown command. Use --help")
	}
}
