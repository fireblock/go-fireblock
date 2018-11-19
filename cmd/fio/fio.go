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
	"time"

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

	verifyHashCmd = app.Command("verify-hash", "verify a sha256 like 0x004e...5c2")
	vhProjectID   = verifyHashCmd.Flag("project-id", "Set project id").Short('p').Default("0x0").String()
	vhUserID      = verifyHashCmd.Flag("user-id", "Set user id").Short('u').Default("").String()
	vhHash        = verifyHashCmd.Arg("file", "File to verify.").Required().String()

	signCmd    = app.Command("sign", "sign a file")
	signKey    = signCmd.Flag("key", "private key in base64url").Short('k').Default("").String()
	signBatch  = signCmd.Flag("batch", "batch name in metadata").Short('b').Default("").String()
	signFio    = signCmd.Flag("fio", "path to fio file").Short('f').ExistingFile()
	passphrase = signCmd.Flag("passphrase", "passphrase (for PGP private key)").Short('p').Default("").String()
	fsign      = signCmd.Arg("file", "File to sign.").Required().ExistingFilesOrDirs()

	versionCmd = app.Command("version", "version of fio (https://fireblock.io)")
)

func exitSuccess(data interface{}, msg string) {
	store := fireblocklib.GetStore("default")
	if verbose := store.GetBool("verbose", false); verbose {
		export, _ := json.Marshal(data)
		fmt.Printf("%s\n", export)
		os.Exit(0)
	} else {
		fmt.Printf("%s\n", msg)
		os.Exit(0)
	}
}

func exitJSONError(data interface{}, code int, msg string) {
	store := fireblocklib.GetStore("default")
	if verbose := store.GetBool("verbose", false); verbose {
		export, _ := json.Marshal(data)
		fmt.Printf("%s\n", export)
		os.Exit(code)
	} else {
		fmt.Printf("%s\n", msg)
		os.Exit(code)
	}
}

// ErrorReturn error return struct
type ErrorReturn struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func exitMsgError(code int, msg string) {
	store := fireblocklib.GetStore("default")
	if verbose := store.GetBool("verbose", false); verbose {
		var data ErrorReturn
		data.Code = code
		data.Message = msg
		export, _ := json.Marshal(data)
		fmt.Printf("%s\n", export)
		os.Exit(code)
	} else {
		fmt.Printf("%s\n", msg)
		os.Exit(code)
	}
}

func exitError(err error) {
	if e, ok := err.(*fireblocklib.FBKError); ok {
		exitMsgError(e.Type(), e.Error())
	} else {
		exitMsgError(fireblocklib.UnknownError, err.Error())
	}
}

// SignSuccess return struct
type SignSuccess struct {
	Code  int      `json:"code"`
	Files []string `json:"files,omitempty"`
}

func signFunction() {
	// Read the private key (signFio or signKey)
	var keyuid, privkey string
	var err error
	// check fio file
	if (signFio == nil || *signFio == "") && (*signKey == "") {
		exitMsgError(fireblocklib.NoFile, "Missing private key! add --fio filepath or --key privatekey")
	} else if signFio != nil && len(*signFio) > 1 {
		keypath := *signFio
		// load fio file
		_, keyuid, privkey, _, err = fireblocklib.LoadFioFile(keypath)
		if err != nil {
			exitMsgError(fireblocklib.InvalidFile, "Invalid fio file")
		}
	} else if signKey != nil && len(*signKey) > 1 {
		keyuid, privkey, err = fireblocklib.LoadB64U(*signKey)
	} else {
		exitMsgError(fireblocklib.NoFile, "Missing private key! add --fio filepath or --key privatekey with correct values")
	}

	// analyze args
	data := make(map[string]fireblocklib.Metadata)
	filepaths := *fsign
	// list filepath
	var filesCertified []string
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
				filesCertified = append(filesCertified, f)
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
			filesCertified = append(filesCertified, filepath)
			data[sha256] = metadata
		}
	}
	if len(data) <= 0 {
		exitMsgError(fireblocklib.NoFile, "Nothing to sign")
	} else if len(data) > 128 {
		exitMsgError(fireblocklib.NoFile, "Batch limitted to 128 objects")
	} else if len(data) == 1 {
		// one certificat
		for hash, metadata := range data {
			signACertificate("", hash, keyuid, privkey, metadata)
			var res SignSuccess
			res.Code = 0
			res.Files = filesCertified
			exitSuccess(res, fmt.Sprintf("File(s) certified: %d", len(filesCertified)))
			return
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
		fname := *signBatch
		if fname == "" {
			t := time.Now()
			y := t.Year()
			mon := int(t.Month())
			d := t.Day()
			h := t.Hour()
			m := t.Minute()
			s := t.Second()
			fname = fmt.Sprintf("batch_%d%d%d_%d:%d:%d", y, mon, d, h, m, s)
		}
		metadata := fireblocklib.Metadata{Kind: "b100", Filename: fname, Size: int64(len(batch)), Type: "application/json"}
		signACertificate(batch, hash, keyuid, privkey, metadata)
		var res SignSuccess
		res.Code = 0
		res.Files = filesCertified
		exitSuccess(res, fmt.Sprintf("File(s) certified: %d", len(filesCertified)))
		return
	}
}

func verifyFunction() {
	if fverify == nil {
		exitMsgError(fireblocklib.NoFile, "Missing file")
	}
	filepath := *fverify
	filename := path.Base(filepath)
	sha256, err := fireblocklib.Sha256File(filepath)
	if err != nil {
		exitMsgError(fireblocklib.InvalidHash, fmt.Sprintf("Cannot compute sha256 on %s", filepath))
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

func verifyHashFunction() {
	sha256 := *vhHash
	if !fireblocklib.IsSha256(sha256) {
		exitMsgError(fireblocklib.InvalidHash, "Invalid Hash")
	}
	if *vhUserID != "" {
		// verify by useruid
		userVerify(*server, "", sha256, *vhUserID, *jverbose)
	} else if *vhProjectID != "0x0" {
		// verify by cardId
		projectVerify(*server, "", sha256, *vhProjectID, *jverbose)
	} else {
		// verify with the first key to register that hash
		verify(*server, "", sha256, *jverbose)
	}
}

// VersionReturn success data
type VersionReturn struct {
	Version string `json:"version,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

func version() {
	var res VersionReturn
	res.Code = 0
	res.Version = Version
	res.Message = fmt.Sprintf("fio %s ( info at https://fireblock.io )", Version)
	exitSuccess(res, res.Message)
}

func main() {
	// check arguments
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate)
	app.Version(Version).Author(Author)
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	// read parsing
	store := fireblocklib.NewStore("default")
	store.SetBool("verbose", *jverbose)
	store.SetString("serverURL", *server)
	// do the command
	switch cmd {
	case signCmd.FullCommand():
		signFunction()
	case verifyCmd.FullCommand():
		verifyFunction()
	case verifyHashCmd.FullCommand():
		verifyHashFunction()
	case versionCmd.FullCommand():
		version()
	}
}
