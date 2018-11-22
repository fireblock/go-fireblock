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

// UserVerifyError project verify
type UserVerifyError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Filename string `json:"filename,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Useruid  string `json:"useruid,omitempty"`
}

func userVerify(server, filename, hash string, useruid string, verbose bool) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/verify-by-user")

	// json inputs + request
	req := UserVerifyReq{hash, useruid}
	res, efbk := Post(url, req)
	if efbk != nil {
		se := UserVerifyError{efbk.Type(), efbk.Error(), filename, hash, useruid}
		exitJSONError(se, efbk.Type(), efbk.Error())
	}

	// parse output
	var response UserVerifyValueReturn
	err := json.Unmarshal(res, &response)
	if err != nil {
		se := UserVerifyError{fireblocklib.InvalidEncoding, err.Error(), filename, hash, useruid}
		exitJSONError(se, fireblocklib.InvalidEncoding, err.Error())
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
		exitSuccess(res, fmt.Sprintf("File %s has been certified by user %s", filename, useruid))
	} else {
		se := UserVerifyError{fireblocklib.InvalidFile, "file not certified on fireblock", filename, hash, useruid}
		exitJSONError(se, fireblocklib.InvalidFile, "file not certified on fireblock")
	}
}
