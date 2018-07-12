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

// ProjectVerifyValue http request
type ProjectVerifyValue struct {
	KeyUID               string `json:"keyUID"`
	Metadata             string `json:"metadata"`
	PkeySignature        string `json:"pkeySignature"`
	CertificateSignature string `json:"certificateSignature"`
	Date                 int64  `json:"date"`
	Pubkey               string `json:"pubkey"`
	KType                string `json:"ktype"`
}

// ProjectVerifyValueReturn http request
type ProjectVerifyValueReturn struct {
	Value []ProjectVerifyValue `json:"value"`
}

// ProjectVerifyResponse http request
type ProjectVerifyResponse struct {
	ID     string                   `json:"id"`
	Errors []common.ErrorRes        `json:"errors,omitempty"`
	Data   ProjectVerifyValueReturn `json:"data"`
}

// ProjectVerifyReq http request
type ProjectVerifyReq struct {
	Hash       string `json:"hash"`
	ProjectUID string `json:"projectuid"`
}

func projectVerify(server, filename, hash string, project *Project, verbose bool) {
	// json inputs
	req := ProjectVerifyReq{hash, project.ProjectUID}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	// http request
	url := "https://$#$server$#$/api//verify-by-project"
	url = strings.Replace(url, "$#$server$#$", server, 1)
	res, err := http.Post(url, "application/json; charset=utf-8", buffer)
	if err != nil {
		verifyError(project, common.NetworkError, fmt.Sprintf("http error %s", url), verbose)
		return
	}
	// analyze result
	var response ProjectVerifyResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		verifyError(project, common.NetworkError, fmt.Sprintf("http response error %s", url), verbose)
		return
	}
	// check result
	if len(response.Errors) > 0 {
		verifyError(project, common.InvalidProject, fmt.Sprintf("Project Error: %s %s", response.Errors[0].ID, response.Errors[0].Detail), verbose)
	}
	// check certificate signature
	validity := false
	values := response.Data.Value
	for _, value := range values {
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
		verifySuccess(project, filename, hash, verbose)
		os.Exit(0)
	} else {
		verifyError(project, common.InvalidFile, fmt.Sprintf("Not a valid file"), verbose)
		os.Exit(1)
	}
}
