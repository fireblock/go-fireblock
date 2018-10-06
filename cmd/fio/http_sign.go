package main

// HTTPSign sign
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

// SignReq http request struct
type SignReq struct {
	Hash              string `json:"hash"`
	KType             string `json:"ktype"`
	Keyuid            string `json:"keyuid"`
	Signature         string `json:"signature"`
	Metadata          string `json:"metadata"`
	MetadataSignature string `json:"metadataSignature"`
}

// CreateCertificateValueReturn http request struct
type CreateCertificateValueReturn struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func createCertificate(server, hash, ktype, keyuid, signature, metadata, metadataSignature string) (string, error) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/create-certificate")

	// json inputs + request
	sig64 := base64.StdEncoding.EncodeToString([]byte(signature))
	meta64 := base64.StdEncoding.EncodeToString([]byte(metadata))
	metaSig64 := base64.StdEncoding.EncodeToString([]byte(metadataSignature))
	req := SignReq{hash, ktype, keyuid, sig64, meta64, metaSig64}
	res, err := Post(url, req)
	if err != nil {
		fbkError(err, false)
		os.Exit(1)
	}

	// parse output
	var response CreateCertificateValueReturn
	err = json.Unmarshal(res, &response)
	if err != nil {
		j, _ := json.Marshal(&res)
		fmt.Println(string(j))
		os.Exit(0)
	}

	return "success", nil

}
