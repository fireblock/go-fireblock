package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fireblock/go-fireblock/common"
)

// GetProjectValueKey http request
type GetProjectValueKey struct {
	Status int64  `json:"status"`
	Pubkey string `json:"pubkey"`
}

// GetProjectValueCard http request
type GetProjectValueCard struct {
	CardUID   string `json:"carduid"`
	Card      string `json:"card"`
	Status    bool   `json:"status,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Date      string `json:"date"`
	Signature string `json:"signature"`
}

// GetProjectValue http request
type GetProjectValue struct {
	Suuid string              `json:"suuid"`
	Key   GetProjectValueKey  `json:"key"`
	Card  GetProjectValueCard `json:"card"`
}

// GetProjectValueReturn http request
type GetProjectValueReturn struct {
	Value GetProjectValue `json:"value"`
}

// GetProjectResponse http request
type GetProjectResponse struct {
	ID     string                `json:"id"`
	Errors []common.ErrorRes     `json:"errors,omitempty"`
	Data   GetProjectValueReturn `json:"data"`
}

// Project result
type Project struct {
	UserUID    string
	Card       string
	CardUID    string
	KType      string
	KeyUID     string
	Pubkey     string
	ProjectUID string
}

func getProject(server, projectuid string) (project *Project, err error) {
	// http request
	url := "$#$server$#$/api/project?projectuid=" + projectuid + "&checkcard=false"
	url = strings.Replace(url, "$#$server$#$", server, 1)
	res, err := http.Get(url)
	if err != nil {
		return nil, common.NewFBKError(fmt.Sprintf("http error %s", url), common.NetworkError)
	}
	// analyze result
	var response GetProjectResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, common.NewFBKError(fmt.Sprintf("http response error %s", url), common.NetworkError)
	}
	// check result
	if len(response.Errors) > 0 {
		return nil, common.NewFBKError(fmt.Sprintf("Project Error: %s %s", response.Errors[0].ID, response.Errors[0].Detail), common.InvalidProject)
	}
	//check cardUID
	card := response.Data.Value.Card.Card
	cardUID := common.Keccak256(card)
	if cardUID != response.Data.Value.Card.CardUID {
		return nil, common.NewFBKError(fmt.Sprintf("Project Error: invalid cardUID"), common.InvalidProject)
	}
	// check project Key
	keyStatus := response.Data.Value.Key.Status
	if keyStatus != 3 {
		return nil, common.NewFBKError(fmt.Sprintf("Project Error: invalid project key state %d", keyStatus), common.InvalidProject)
	}
	pubkey := response.Data.Value.Key.Pubkey
	userUID := ""
	KType := "unknown"
	KeyUID := "0x0"
	// check card signature
	if cardUID != "0x0" {
		cardSignature := response.Data.Value.Card.Signature
		date := response.Data.Value.Card.Date
		message := fmt.Sprintf("register card %s at %s", cardUID, date)
		sigCheck, err2 := common.ECDSAVerify(pubkey, message, cardSignature)
		if err2 != nil {
			return nil, common.NewFBKError(fmt.Sprintf("Project Error: invalid project bad signature"), common.InvalidProject)
		}
		if !sigCheck {
			return nil, common.NewFBKError(fmt.Sprintf("Project Error: invalid project bad signature card"), common.InvalidProject)
		}
		// check card
		useruid, keyuid, ktype, err3 := common.CheckCard(card, cardUID)
		if err3 != nil {
			return nil, common.NewFBKError(fmt.Sprintf("Project Error: invalid project bad card: %s", err3), common.InvalidProject)
		}
		userUID = useruid
		KType = ktype
		KeyUID = keyuid
	}
	// project
	project = new(Project)
	project.ProjectUID = projectuid
	project.Card = card
	project.CardUID = cardUID
	project.KType = KType
	project.KeyUID = KeyUID
	project.Pubkey = pubkey
	project.UserUID = userUID
	return project, nil
}
