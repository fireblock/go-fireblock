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
	"encoding/json"
	"fmt"

	"github.com/fireblock/go-fireblock/fireblocklib"
)

// SignReq http request struct
type SignReq struct {
	Batch             string `json:"batch,omitempty"`
	Hash              string `json:"hash"`
	KType             string `json:"ktype"`
	Keyuid            string `json:"keyuid"`
	Signature         string `json:"signature"`
	Metadata          string `json:"metadata"`
	MetadataSignature string `json:"metadataSignature"`
}

// Batch batch
type Batch struct {
	Kind     string `json:"kind"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
}

// CreateCertificateValueReturn http request struct
type CreateCertificateValueReturn struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// SignError signature error
type SignError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Batch   string `json:"batch,omitempty"`
	Hash    string `json:"hash,omitempty"`
	KType   string `json:"ktype,omitempty"`
	Keyuid  string `json:"keyuid,omitempty"`
}

func createCertificate(server, hash, ktype, keyuid, signature, metadata, metadataSignature string) (string, error) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/create-certificate")

	// json inputs + request
	req := SignReq{"", hash, ktype, keyuid, signature, metadata, metadataSignature}
	res, efbk := Post(url, req)
	if efbk != nil {
		se := SignError{efbk.Type(), efbk.Error(), "", hash, ktype, keyuid}
		exitJSONError(se, efbk.Type(), efbk.Error())
	}

	// parse output
	var response CreateCertificateValueReturn
	err := json.Unmarshal(res, &response)
	if err != nil {
		se := SignError{fireblocklib.InvalidEncoding, err.Error(), "", hash, ktype, keyuid}
		exitJSONError(se, fireblocklib.InvalidEncoding, err.Error())
	}

	return "success", nil
}

func createBatch(server, batch, hash, ktype, keyuid, signature, metadata, metadataSignature string) (string, error) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/create-batch")

	// json inputs + request
	req := SignReq{batch, hash, ktype, keyuid, signature, metadata, metadataSignature}
	res, efbk := Post(url, req)
	if efbk != nil {
		se := SignError{efbk.Type(), efbk.Error(), batch, hash, ktype, keyuid}
		exitJSONError(se, efbk.Type(), efbk.Error())
	}

	// parse output
	var response CreateCertificateValueReturn
	err := json.Unmarshal(res, &response)
	if err != nil {
		se := SignError{fireblocklib.InvalidEncoding, err.Error(), batch, hash, ktype, keyuid}
		exitJSONError(se, fireblocklib.InvalidEncoding, err.Error())
	}

	return "success", nil
}

func signACertificate(batch, hash, keyuid, privkey string, metadata string) {
	// create message
	message := hash + "||" + keyuid
	messageSignature := ""
	if metadata != "" {
		metadataSID := fireblocklib.Keccak256(metadata)
		messageSignature = metadataSID + "||" + hash + "||" + keyuid
	}
	// create signature
	var err error
	signature := ""
	metadataSignature := ""
	ktype := fireblocklib.B32Type(keyuid)
	if ktype == "pgp" {
		signature, err = fireblocklib.PGPSign(privkey, message, *passphrase)
		if err != nil {
			se := SignError{fireblocklib.SignError, "PGP error: cannot sign message", batch, hash, ktype, keyuid}
			exitJSONError(se, fireblocklib.InvalidEncoding, "PGP error: cannot sign message")
		}
		if metadata != "" {
			metadataSignature, err = fireblocklib.PGPSign(privkey, messageSignature, *passphrase)
			if err != nil {
				se := SignError{fireblocklib.SignError, "PGP error: cannot sign metadata", batch, hash, ktype, keyuid}
				exitJSONError(se, fireblocklib.InvalidEncoding, "PGP error: cannot sign metadata")
			}
		}
	} else if ktype == "ecdsa" {
		signature, err = fireblocklib.ECDSASign(privkey, message)
		if err != nil {
			se := SignError{fireblocklib.SignError, "ECDSA error: cannot sign message", batch, hash, ktype, keyuid}
			exitJSONError(se, fireblocklib.InvalidEncoding, "ECDSA error: cannot sign message")
		}
		if metadata != "" {
			metadataSignature, err = fireblocklib.ECDSASign(privkey, messageSignature)
			if err != nil {
				se := SignError{fireblocklib.SignError, "ECDSA error: cannot sign metadata", batch, hash, ktype, keyuid}
				exitJSONError(se, fireblocklib.InvalidEncoding, "ECDSA error: cannot sign metadata")
			}
		}
	} else {
		msg := fmt.Sprintf("Invalid key format %s\n", ktype)
		se := SignError{fireblocklib.SignError, msg, batch, hash, ktype, keyuid}
		exitJSONError(se, fireblocklib.InvalidEncoding, msg)
	}
	// sign
	if batch == "" {
		_, err = createCertificate(*server, hash, ktype, keyuid, signature, metadata, metadataSignature)
	} else {
		_, err = createBatch(*server, batch, hash, ktype, keyuid, signature, metadata, metadataSignature)
	}
	if err != nil {
		fmt.Println(err)
		se := SignError{fireblocklib.SignError, "unknown error", batch, hash, ktype, keyuid}
		exitJSONError(se, fireblocklib.InvalidEncoding, "unknown error")
	}
}
