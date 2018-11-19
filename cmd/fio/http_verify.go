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

// VerifyValue http request
type VerifyValue struct {
	PKey        KeyInfo         `json:"pkey"`
	Card        CardInfo        `json:"card"`
	Key         KeyInfo         `json:"key"`
	Certificate CertificateInfo `json:"certificate"`
}

// VerifyValueReturn http request
type VerifyValueReturn struct {
	Value []VerifyValue `json:"value"`
}

// VerifyReq http request
type VerifyReq struct {
	Hash string `json:"hash"`
}

func verify(server, filename, hash string, verbose bool) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/verify-by-hash")

	// json inputs + request
	req := VerifyReq{hash}
	res, err := Post(url, req)
	if err != nil {
		exitError(err)
	}

	// parse output
	var response VerifyValueReturn
	err = json.Unmarshal(res, &response)
	if err != nil {
		exitError(err)
	}

	// check certificate signature
	var pkey, key KeyInfo
	var card CardInfo
	var certificate CertificateInfo
	validity := false
	values := response.Value
	for _, value := range values {

		pkey = value.PKey
		key = value.Key
		card = value.Card
		certificate = value.Certificate

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
		exitMsgError(fireblocklib.InvalidFile, fmt.Sprintf("Not a valid file"))
	}
}
