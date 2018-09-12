package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fireblock/go-fireblock/common"
)

// GetProjectValueCard http request
type GetProjectValueCard struct {
	UID       string `json:"uid"`
	Txt       string `json:"txt"`
	Date      int64  `json:"date"`
	Signature string `json:"signature"`
}

// GetProjectValueKey http request
type GetProjectValuePKey struct {
	Keyuid string              `json:"keyuid"`
	KType  string              `json:"keytype"`
	State  int64               `json:"state"`
	Pubkey string              `json:"pubkey"`
	Card   GetProjectValueCard `json:"card"`
}

type GetProjectValueUser struct {
	Suuid string `json:"suuid"`
}

// GetProjectValue http request
type GetProjectValue struct {
	User GetProjectValueUser `json:"user"`
	PKey GetProjectValuePKey `json:"pkey"`
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

func convertPS2CIP(ps common.ProviderState) (cip CardInfoProvider) {
	cip.UID = ps.UID
	cip.Status = ps.Status
	cip.Proof = ps.Proof
	return cip
}

func getProject(server, projectuid string) (pi ProjectInfo, ci CardInfo, err error) {
	// http request
	url := "$#$server$#$/api/project?projectuid=" + projectuid
	url = strings.Replace(url, "$#$server$#$", server, 1)
	res, err := http.Get(url)
	if err != nil {
		return ProjectInfo{}, CardInfo{}, common.NewFBKError(fmt.Sprintf("http error %s", url), common.NetworkError)
	}
	// analyze result
	var response GetProjectResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return ProjectInfo{}, CardInfo{}, common.NewFBKError(fmt.Sprintf("http response error %s", url), common.NetworkError)
	}
	// check result
	if len(response.Errors) > 0 {
		return ProjectInfo{}, CardInfo{}, common.NewFBKError(fmt.Sprintf("Project Error: %s %s", response.Errors[0].ID, response.Errors[0].Detail), common.InvalidProject)
	}
	// extract projectInfo
	suuid := response.Data.Value.User.Suuid
	pkey := response.Data.Value.PKey
	pi = ProjectInfo{}
	pi.UID = pkey.Keyuid
	pi.KType = pkey.KType
	pi.State = pkey.State
	pi.Pubkey = pkey.Pubkey
	if (pi.State & 7) == 3 {
		pi.Status = "ok"
	} else {
		return ProjectInfo{}, CardInfo{}, common.NewFBKError(fmt.Sprintf("Project Error: invalid project key state"), common.InvalidProject)
	}
	// extract cardInfo
	card := response.Data.Value.PKey.Card
	ci = CardInfo{}
	if card.Txt == "" {
		ci.Status = "none"
		ci.Txt = ""
		ci.UID = "0x0"
	} else {
		// check cardUID
		carduid := common.Keccak256(card.Txt)
		if carduid != card.UID {
			ci.Status = "error"
			return ProjectInfo{}, CardInfo{}, common.NewFBKError(fmt.Sprintf("Project Error: invalid cardUID"), common.InvalidProject)
		}
		// check signature of the card
		msg := fmt.Sprintf("register card %s at %d", carduid, card.Date)
		ck, err := common.ECDSAVerify(pi.Pubkey, msg, card.Signature)
		if err != nil || !ck {
			return ProjectInfo{}, CardInfo{}, common.NewFBKError(fmt.Sprintf("Project Error: invalid signature of the card"), common.InvalidProject)
		}
		ci.Txt = card.Txt
		ci.Signature = card.Signature
		ci.UID = card.UID
		ci.Date = card.Date
		// check card
		providers, err3 := common.VerifyCard(ci.Txt, pi.UID, pi.KType, suuid)
		if err3 != nil {
			return ProjectInfo{}, CardInfo{}, err3
		}
		ci.Twitter = convertPS2CIP(providers.Twitter)
		ci.Github = convertPS2CIP(providers.Github)
		ci.Linkedin = convertPS2CIP(providers.Linkedin)
		ci.Https = convertPS2CIP(providers.Https)
		ci.Status = "ok"
	}
	return pi, ci, nil
}
