// Copyright 2015-2017 Fireblock.
// This file is part of Fireblock.

// Fireblock is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Fireblock is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Fireblock.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/fireblock/go-fireblock/fireblocklib"
)

func convertPS2CIP(ps fireblocklib.ProviderState) (cip CardInfoProvider) {
	cip.UID = ps.UID
	cip.Status = ps.Status
	cip.Proof = ps.Proof
	return cip
}

// ProjectVerifyResult result
type ProjectVerifyResult struct {
	Key         KeyInfo         `json:"key"`
	Certificate CertificateInfo `json:"certificate"`
	Card        CardInfo        `json:"card"`
	PKey        KeyInfo         `json:"pkey"`
}

// ProjectVerifyValue http request
// keyUID, metadata, pkeySignature, certificateSignature, date, cdate, pkeystate: pks, ppubkey, keystate: ks, pubkey, ktype, metadataSignature
type ProjectVerifyValue struct {
	Results []ProjectVerifyResult `json:"results"`
}

// ProjectVerifyValueReturn http request
type ProjectVerifyValueReturn struct {
	Value []ProjectVerifyResult `json:"value"`
	ID    string                `json:"id"`
}

// ProjectVerifyReq http request
type ProjectVerifyReq struct {
	Hash       string `json:"hash"`
	ProjectUID string `json:"projectuid"`
}

// ProjectVerifySuccess result
type ProjectVerifySuccess struct {
	Filename    string          `json:"filename"`
	Key         KeyInfo         `json:"key"`
	Certificate CertificateInfo `json:"certificate"`
	Card        CardInfo        `json:"card"`
	PKey        KeyInfo         `json:"pkey"`
}

func projectVerify(server, filename, hash, pkeyUID string, verbose bool) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/verify-by-project")

	// json inputs + request
	req := ProjectVerifyReq{hash, pkeyUID}
	res, err := Post(url, req)
	if err != nil {
		exitError(err)
	}

	// parse output
	var response ProjectVerifyValueReturn
	err = json.Unmarshal(res, &response)
	if err != nil {
		exitError(err)
	}

	var pkey, key KeyInfo
	var card CardInfo
	var certificate CertificateInfo
	validity := false

	values := response.Value
	for _, value := range values {
		key = value.Key
		certificate = value.Certificate
		pkey = value.PKey
		card = value.Card

		ck := checkAResult(pkey, key, &card, certificate, hash)
		if !ck {
			continue
		}
		validity = true
		break
	}
	if validity {
		var res ProjectVerifySuccess
		res.Filename = filename
		res.Card = card
		res.Certificate = certificate
		res.PKey = pkey
		res.Key = key
		exitSuccess(res, fmt.Sprintf("File %s has been certified by project %s", filename, pkey.KeyUID))
	} else {
		if pkey.KeyUID == "" {
			exitMsgError(fireblocklib.InvalidProject, "Invalid project")
		}
		exitMsgError(fireblocklib.InvalidFile, fmt.Sprintf("file not certified on fireblock"))
	}
}
