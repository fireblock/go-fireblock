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
	"encoding/json"
	"fmt"
	"os"
	"path"

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
	fsign      = signCmd.Arg("file", "File to sign.").Required().ExistingFilesOrDirs()

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

	// analyze args
	data := make(map[string]fireblocklib.Metadata)
	filepaths := *fsign
	for _, filepath := range filepaths {
		// check if directory or file
		stat, err := os.Stat(filepath)
		if err != nil {
			// skip
			continue
		}
		if stat.Mode().IsDir() {
			files, err := fireblocklib.ListFilesInDirectory(filepath)
			if err != nil {
				fmt.Printf("Cannot list files in %s\n", filepath)
				continue
			}
			for _, f := range files {
				sha256, err := fireblocklib.Sha256File(f)
				if err != nil {
					fmt.Printf("Cannot compute sha256 on %s\n", f)
					continue
				}
				// get metadata
				metadata, err := fireblocklib.MetadataFile(f)
				if err != nil {
					fmt.Printf("Cannot read metadata on %s\n", f)
					continue
				}
				data[sha256] = metadata
			}
		}
		if stat.Mode().IsRegular() {
			sha256, err := fireblocklib.Sha256File(filepath)
			if err != nil {
				fmt.Printf("Cannot compute sha256 on %s\n", filepath)
				continue
			}
			// get metadata
			metadata, err := fireblocklib.MetadataFile(filepath)
			if err != nil {
				fmt.Printf("Cannot read metadata on %s\n", filepath)
				continue
			}
			data[sha256] = metadata
		}
	}
	if len(data) <= 0 {
		fmt.Printf("Nothing to sign\n")
		os.Exit(1)
	} else if len(data) > 128 {
		fmt.Printf("Batch limitted to 128 objects\n")
		os.Exit(1)
	} else if len(data) == 1 {
		// one certificat
		for hash, metadata := range data {
			signACertificate("", hash, keyuid, privkey, metadata)
			break
		}
	} else {
		// a batch
		var batchArray []fireblocklib.BatchElem
		for hash, metadata := range data {
			batchArray = append(batchArray, fireblocklib.BatchElem{Hash: hash, Filename: metadata.Filename, Size: metadata.Size, Type: metadata.Type})
		}
		b, _ := json.Marshal(&batchArray)
		batch := string(b)
		// compute sha256
		hash := fireblocklib.Sha256(batch)
		metadata := fireblocklib.Metadata{Kind: "b100", Filename: "batch", Size: int64(len(batch)), Type: "application/json"}
		signACertificate(batch, hash, keyuid, privkey, metadata)
	}
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
	fmt.Printf("fio %s ( info at https://fireblock.io )\n", Version)
	os.Exit(0)
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate)
	app.Version(Version).Author(Author)
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
