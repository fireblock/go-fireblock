package main

import (
	"fmt"
	"os"
	"path"

	"github.com/fireblock/go-fireblock/common"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	// Version - fbk version
	Version = "0.1.2"
	Author  = "Laurent Mallet laurent.mallet at gmail dot com"
)

func main() {
	cardId := kingpin.Flag("card-id", "Set card id").Short('c').Default("0x0").String()
	userId := kingpin.Flag("user-id", "Set user id").Short('u').Default("0x0").String()
	json := kingpin.Flag("json", "Output in JSON format").Short('j').Bool()
	file := kingpin.Arg("file", "File to retrieve.").Required().ExistingFile()

	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(Version).Author(Author)
	kingpin.Parse()
	if file != nil {
		if *cardId == "0x0" && *userId == "0x0" {
			fmt.Printf("Use -u or -c\n")
			os.Exit(1)
		}
		filepath := *file
		filename := path.Base(filepath)
		sha256, err := common.Sha256File(filepath)
		if err != nil {
			fmt.Printf("Cannot compute sha256 on %s\n", filepath)
			os.Exit(1)
		}
		if *userId != "0x0" {
			// verify by useruid
			common.GVerify(filename, sha256, *userId, *json)
		} else {
			// verify by cardId
			common.CVerify(filename, sha256, *cardId, *json)
		}
	}
}
