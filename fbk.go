package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fireblock/go-fireblock/common"
)

const (
	// Version - fbk version
	Version = "0.0.1"
)

func version() {
	fmt.Printf("FBK v%v\n", Version)
}

func verifyByUseruid() {

}

func verifyByCardId() {

}

func main() {
	isVersion := flag.Bool("version", false, "display package version")
	hasUseruidPtr := flag.String("useruid", "", "verify by useruid")
	hasCardIdPtr := flag.String("cardId", "", "verify by cardId")
	flag.Parse()
	// Print versions
	if *isVersion {
		version()
		return
	}
	args := flag.Args()
	if len(args) != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	filepath := args[0]
	// check file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Printf("File %s doesn't exist or readable\n", filepath)
		os.Exit(1)
	}
	// check there is a useruid or cardId provided
	if *hasUseruidPtr == "" && *hasCardIdPtr == "" {
		fmt.Printf("Use -useruid or -cardId option\n")
		os.Exit(1)
	}
	// check only option is selected
	if *hasUseruidPtr != "" && *hasCardIdPtr != "" {
		fmt.Printf("Use only one option -useruid or -cardId option\n")
		os.Exit(1)
	}
	sha256, err := common.Sha256File(filepath)
	if err != nil {
		fmt.Printf("Cannot compute sha256 on %s", filepath)
		os.Exit(1)
	}
	if *hasUseruidPtr != "" {
		// verify by useruid
		common.GVerify(sha256, *hasUseruidPtr)
		fmt.Println(sha256)
	} else {
		// verify by cardId
		sha256, _ := common.Sha256File(filepath)
		fmt.Println(sha256)
	}
}
