package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fireblock/go-fireblock/contracts"
)

// GetKey get public key
func GetKey(store *contracts.Store, keyUID string) (struct {
	Key     string
	Revoked bool
	Begin   int64
	End     int64
}, error) {
	b, _ := hexutil.Decode(keyUID)
	hash := common.BytesToHash(b)
	key, _ := store.GetKey(nil, hash)
	return key, nil
}

func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("/home/ellis/.ethereum/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	store, err := contracts.NewStore(common.HexToAddress("0x0cFD5bD8889eef4d352A12C794dA3d5Ec64Ab8BD"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	if store != nil {
		fmt.Printf("")
	}
	// b, err := hexutil.Decode("0x10000000000000000000000074da8b98e4d2471f7cbc39c27caf0d730dcec923")
	// hash := common.BytesToHash(b)
	// key, err := store.GetKey(nil, hash)
	key, _ := GetKey(store, "0x10000000000000000000000074da8b98e4d2471f7cbc39c27caf0d730dcec923")
	if err != nil {
		log.Fatalf("Failed to retrieve pub key: %v", err)
	}
	fmt.Println("pub key:", key.Key)
	fmt.Println("pub key:", key.Revoked)
	fmt.Println("pub key:", key.Begin)
	fmt.Println("pub key:", key.End)
}
