package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	fireblock "github.com/fireblock/go-fireblock"
	fcommon "github.com/fireblock/go-fireblock/common"
	"github.com/fireblock/go-fireblock/contracts"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func checkFilePath(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func getConn(ipcPath, rpcPath, wsPath *string) *ethclient.Client {
	connMode := 0
	var path *string
	if *ipcPath != "" {
		connMode++
		// check file exist
		if !checkFilePath(*ipcPath) {
			fmt.Printf("No ipc file %s", *ipcPath)
			os.Exit(1)
		}
		path = ipcPath
	}
	if *rpcPath != "" {
		connMode++
		path = rpcPath
	}
	if *wsPath != "" {
		connMode++
		path = wsPath
	}
	if connMode == 0 {
		fmt.Printf("Set a ipc/rpc/ws connection")
		os.Exit(1)
	}
	if connMode > 1 {
		fmt.Printf("Only one of ipc/rpc/ws allowed")
		os.Exit(1)
	}
	// connexion by ipc
	var conn *ethclient.Client
	var err error
	conn, err = ethclient.Dial(*path)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	block, err := conn.BlockByNumber(ctx, nil)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: Can't retrieve last block")
		os.Exit(1)
	}
	fmt.Printf("%d", block.Number())
	return conn
}

func main() {
	cardID := kingpin.Flag("card-id", "Set card id").Short('c').Default("0x0").String()
	userID := kingpin.Flag("user-id", "Set user id").Short('u').Default("0x0").String()
	json := kingpin.Flag("json", "Output in JSON format").Short('j').Bool()
	file := kingpin.Arg("file", "File to retrieve.").Required().ExistingFile()

	ipcPath := kingpin.Flag("ipc", "ipc path").Default("").String()
	rpcPath := kingpin.Flag("rpc", "rpc url").Default("").String()
	wsPath := kingpin.Flag("ws", "ws url").Default("").String()

	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(fireblock.Version).Author(fireblock.Author)
	kingpin.Parse()

	// Create an IPC based RPC connection to a remote node
	conn := getConn(ipcPath, rpcPath, wsPath)

	// Instantiate the store contract
	store, err := contracts.NewStore(common.HexToAddress("0x0cFD5bD8889eef4d352A12C794dA3d5Ec64Ab8BD"), conn)
	if err != nil {
		fmt.Printf("Failed to instantiate the store contract: %v", err)
		os.Exit(1)
	}
	// Instantiate the fireblock contract
	fireblock, err := contracts.NewFireblock(common.HexToAddress("0x9202DAe4De513e73c594E81a1Fd7eaCc99B470a2"), conn)
	if err != nil {
		fmt.Printf("Failed to instantiate the fireblock contract: %v\n", err)
		os.Exit(1)
	}

	// check if there's a file
	if file == nil {
		fmt.Printf("No file found")
		os.Exit(1)
	}
	if *cardID == "0x0" && *userID == "0x0" {
		fmt.Printf("Use -u or -c\n")
		os.Exit(1)
	}

	// compute sha256
	filepath := *file
	filename := path.Base(filepath)
	sha256, err := fcommon.Sha256File(filepath)
	if err != nil {
		fmt.Printf("Cannot compute sha256 on %s\n", filepath)
		os.Exit(1)
	}

	if *userID != "0x0" {
		// verify by useruid
		fmt.Print(filename)
		fmt.Print(json)
		VerifyByUser(fireblock, store, filename, sha256, *userID, *json)
	} else {
		// verify by cardId
		fmt.Print(sha256)
		Verify(fireblock, store, filename, sha256, *cardID, *json)
	}
}
