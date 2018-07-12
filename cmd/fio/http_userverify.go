package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fireblock/go-fireblock/common"
)

// UserVerifyValue http request
type UserVerifyValue struct {
	PKeyUID              string `json:"pkeyUID"`
	KeyUID               string `json:"keyUID"`
	Metadata             string `json:"metadata"`
	PkeySignature        string `json:"pkeySignature"`
	CertificateSignature string `json:"certificateSignature"`
	Date                 int64  `json:"date"`
	Pubkey               string `json:"pubkey"`
	KType                string `json:"ktype"`
}

// UserVerifyValueReturn http request
type UserVerifyValueReturn struct {
	Value []UserVerifyValue `json:"value"`
}

// UserVerifyResponse http request
type UserVerifyResponse struct {
	ID     string                `json:"id"`
	Errors []common.ErrorRes     `json:"errors,omitempty"`
	Data   UserVerifyValueReturn `json:"data"`
}

// UserVerifyReq http request
type UserVerifyReq struct {
	Hash    string `json:"hash"`
	UserUID string `json:"useruuid"`
}

func userVerify(server, filename, hash string, useruid string, verbose bool) (project2 *Project, err error) {
	// json inputs
	req := UserVerifyReq{hash, useruid}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	// http request
	url := "https://$#$server$#$/api/verify-by-user"
	url = strings.Replace(url, "$#$server$#$", server, 1)
	res, err := http.Post(url, "application/json; charset=utf-8", buffer)
	if err != nil {
		return nil, common.NewFBKError(fmt.Sprintf("http error %s", url), common.NetworkError)
	}
	// analyze result
	var response UserVerifyResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, common.NewFBKError(fmt.Sprintf("http response error %s", url), common.NetworkError)
	}
	// check result
	if len(response.Errors) > 0 {
		return nil, common.NewFBKError(fmt.Sprintf("Project Error: %s %s", response.Errors[0].ID, response.Errors[0].Detail), common.InvalidProject)
	}
	// check certificate signature
	validity := false
	values := response.Data.Value
	for _, value := range values {
		project, err := getProject(server, value.PKeyUID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// check certificate
		message := fmt.Sprintf("%s||%s", hash, value.KeyUID)
		if value.KType == "ecdsa" {
			ck, err := common.ECDSAVerify(value.Pubkey, message, value.CertificateSignature)
			if err != nil {
				continue
			}
			if !ck {
				continue
			}
		} else if value.KType == "pgp" {
			ck, err := common.PGPVerify(value.Pubkey, message, value.CertificateSignature)
			if err != nil {
				continue
			}
			if !ck {
				continue
			}
		}
		// check delegation
		message2 := fmt.Sprintf("approved key is %s at %d", value.KeyUID, value.Date)
		ck, err := common.ECDSAVerify(project.Pubkey, message2, value.PkeySignature)
		if err != nil {
			continue
		}
		if !ck {
			continue
		}
		validity = true
		break
	}
	if validity {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}
	return nil, nil
}