package main

import (
	"fmt"
	"os"
	"path"

	fireblock "github.com/fireblock/go-fireblock/"
	"github.com/fireblock/go-fireblock/common"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	cardID := kingpin.Flag("card-id", "Set card id").Short('c').Default("0x0").String()
	userID := kingpin.Flag("user-id", "Set user id").Short('u').Default("0x0").String()
	json := kingpin.Flag("json", "Output in JSON format").Short('j').Bool()
	file := kingpin.Arg("file", "File to retrieve.").Required().ExistingFile()

	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(fireblock.Version).Author(fireblock.Author)
	kingpin.Parse()
	if file != nil {
		if *cardID == "0x0" && *userID == "0x0" {
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
		if *userID != "0x0" {
			// verify by useruid
			GVerify(filename, sha256, *userID, *json)
		} else {
			// verify by cardId
			CVerify(filename, sha256, *cardID, *json)
		}
	}
}
