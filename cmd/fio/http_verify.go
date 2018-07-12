package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/fireblock/go-fireblock/common"
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

func verifyError(project *Project, code int, message string, verbose bool) {
	if verbose {
		var r VerifyErrorReturn
		r.Error = strconv.Itoa(code)
		r.Detail = message
		if project != nil {
			r.ProjectUID = project.ProjectUID
		}
		export, _ := json.Marshal(r)
		fmt.Printf("%s\n", export)
	} else {
		fmt.Printf("error %d: %s\n", code, message)
	}
	os.Exit(1)
}

func verifySuccess(project *Project, filename, hash string, verbose bool) {
	if verbose {
		var r VerifySuccessReturn
		r.Verified = true
		r.Filename = filename
		r.Hash = hash
		r.ProjectUID = project.ProjectUID
		r.CardUID = project.CardUID
		r.Card = project.Card
		export, _ := json.Marshal(r)
		fmt.Printf("%s\n", export)
	} else {
		fmt.Printf("VALID FILE\n")
	}
	os.Exit(0)
}

func pVerify(server, filename, hash, projectuid string, verbose bool) {
	project, err := getProject(server, projectuid)
	if err != nil {
		if err, ok := err.(*common.FBKError); ok {
			verifyError(nil, err.Type(), err.Error(), verbose)
			os.Exit(1)
		}
		verifyError(nil, common.UnknownError, err.Error(), verbose)
		os.Exit(1)
	}
	projectVerify(server, filename, hash, project, verbose)
}

func uVerify(server, filename, hash, useruid string, verbose bool) {
	userVerify(server, filename, hash, useruid, verbose)
}
