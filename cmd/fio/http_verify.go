package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// VerifySuccessReturn success data
type VerifySuccessReturn struct {
	Verified   bool   `json:"verified"`
	Hash       string `json:"hash"`
	Filename   string `json:"filename"`
	UserUID    string `json:"useruid,omitempty"`
	CardUID    string `json:"carduid,omitempty"`
	ProjectUID string `json:"projectuid,omitempty"`
	Card       string `json:"card,omitempty"`
}

// VerifyErrorReturn error return struct
type VerifyErrorReturn struct {
	Error      string `json:"error"`
	Detail     string `json:"detail"`
	Hash       string `json:"hash,omitempty"`
	Filename   string `json:"filename,omitempty"`
	UserID     string `json:"useruid,omitempty"`
	CardID     string `json:"carduid,omitempty"`
	ProjectUID string `json:"projectuid,omitempty"`
}

func verifyError(projectInfo KeyInfo, cardInfo CardInfo, code int, message string, verbose bool) {
	if verbose {
		var r VerifyErrorReturn
		r.Error = strconv.Itoa(code)
		r.Detail = message
		if projectInfo.Status != "ok" {
			r.ProjectUID = projectInfo.KeyUID
		}
		export, _ := json.Marshal(r)
		fmt.Printf("%s\n", export)
	} else {
		fmt.Printf("error %d: %s\n", code, message)
	}
	os.Exit(1)
}

func verifySuccess(projectInfo KeyInfo, cardInfo CardInfo, filename, hash string, verbose bool) {
	if verbose {
		var r VerifySuccessReturn
		r.Verified = true
		r.Filename = filename
		r.Hash = hash
		r.ProjectUID = projectInfo.KeyUID
		r.CardUID = cardInfo.UID
		r.Card = cardInfo.Txt
		export, _ := json.Marshal(r)
		fmt.Printf("%s\n", export)
	} else {
		fmt.Printf("VALID FILE\n")
	}
	os.Exit(0)
}
