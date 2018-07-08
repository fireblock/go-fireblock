package main

// HTTPSign sign
import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fireblock/go-fireblock/common"
)

// SignReq http request struct
type SignReq struct {
	Hash      string `json:"hash"`
	Keyuid    string `json:"keyuid"`
	Signature string `json:"signature"`
	Metadata  string `json:"metadata"`
}

// CreateCertificateValueReturn http request struct
type CreateCertificateValueReturn struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// CreateCertificateResponse http request struct
type CreateCertificateResponse struct {
	Errors []common.ErrorRes            `json:"errors,omitempty"`
	Data   CreateCertificateValueReturn `json:"data"`
}

func createCertificate(server, hash, keyuid, signature, metadata string) (string, error) {
	sig64 := base64.StdEncoding.EncodeToString([]byte(signature))
	meta64 := base64.StdEncoding.EncodeToString([]byte(metadata))
	req := SignReq{hash, keyuid, sig64, meta64}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	url := "https://$#$server$#$/api/create-certificate"
	url = strings.Replace(url, "$#$server$#$", server, 1)
	res, err := http.Post(url, "application/json; charset=utf-8", buffer)
	if err != nil {
		return "", common.NewFBKError(fmt.Sprintf("http error %s", url), common.NetworkError)
	}
	var response CreateCertificateResponse
	err2 := json.NewDecoder(res.Body).Decode(&response)
	if err2 != nil {
		return "", common.NewFBKError("", common.UnknownError)
	}
	// check errors in response
	if len(response.Errors) > 0 {
		err := response.Errors[0]
		return "", common.NewFBKError(err.ID, common.InvalidSignature)
	}
	if response.Data.ID == "success" {
		return "success", nil
	}
	return "", common.NewFBKError("", common.UnknownError)
}
