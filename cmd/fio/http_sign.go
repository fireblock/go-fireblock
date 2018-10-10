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

// HTTPSign sign
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

// SignReq http request struct
type SignReq struct {
	Hash              string `json:"hash"`
	KType             string `json:"ktype"`
	Keyuid            string `json:"keyuid"`
	Signature         string `json:"signature"`
	Metadata          string `json:"metadata"`
	MetadataSignature string `json:"metadataSignature"`
}

// CreateCertificateValueReturn http request struct
type CreateCertificateValueReturn struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func createCertificate(server, hash, ktype, keyuid, signature, metadata, metadataSignature string) (string, error) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/create-certificate")

	// json inputs + request
	sig64 := base64.StdEncoding.EncodeToString([]byte(signature))
	meta64 := base64.StdEncoding.EncodeToString([]byte(metadata))
	metaSig64 := base64.StdEncoding.EncodeToString([]byte(metadataSignature))
	req := SignReq{hash, ktype, keyuid, sig64, meta64, metaSig64}
	res, err := Post(url, req)
	if err != nil {
		fbkError(err, false)
		os.Exit(1)
	}

	// parse output
	var response CreateCertificateValueReturn
	err = json.Unmarshal(res, &response)
	if err != nil {
		j, _ := json.Marshal(&res)
		fmt.Println(string(j))
		os.Exit(0)
	}

	return "success", nil

}
