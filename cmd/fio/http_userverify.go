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
	"os"

	"github.com/fireblock/go-fireblock/fireblocklib"
)

// UserVerifyValue http request
type UserVerifyValue struct {
	PKey        KeyInfo         `json:"pkey"`
	Card        CardInfo        `json:"card"`
	Key         KeyInfo         `json:"key"`
	Certificate CertificateInfo `json:"certificate"`
}

// UserVerifyValueReturn http request
type UserVerifyValueReturn struct {
	Value []UserVerifyValue `json:"value"`
}

// UserVerifyResponse http request
type UserVerifyResponse struct {
	ID     string                  `json:"id"`
	Errors []fireblocklib.ErrorRes `json:"errors,omitempty"`
	Data   UserVerifyValueReturn   `json:"data"`
}

// UserVerifyReq http request
type UserVerifyReq struct {
	Hash    string `json:"hash"`
	UserUID string `json:"useruuid"`
}

func userVerify(server, filename, hash string, useruid string, verbose bool) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/verify-by-user")

	// json inputs + request
	req := UserVerifyReq{hash, useruid}
	res, err := Post(url, req)
	if err != nil {
		fbkError(err, verbose)
		os.Exit(1)
	}

	// parse output
	var response UserVerifyValueReturn
	err = json.Unmarshal(res, &response)
	if err != nil {
		j, _ := json.Marshal(&res)
		fmt.Print(string(j))
		os.Exit(1)
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

		ck := checkAResult(pkey, key, card, certificate, hash)
		if !ck {
			continue
		}
		validity = true
		break
	}
	if validity {
		verifySuccess(pkey, card, certificate, filename, hash, verbose)
		os.Exit(0)
	} else {
		verifyError(pkey, card, fireblocklib.InvalidFile, fmt.Sprintf("Not a valid file"), verbose)
		os.Exit(1)
	}
}
