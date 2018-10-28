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

// VerifyValueResponse http request
type VerifyValueResponse struct {
	Value  VerifyValueReturn       `json:"value"`
	Errors []fireblocklib.ErrorRes `json:"errors,omitempty"`
	ID     string                  `json:"id"`
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
		fbkError(err, verbose)
		os.Exit(1)
	}

	// parse output
	var response VerifyValueReturn
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

		// pkey state
		if (pkey.State & 15) != 3 {
			continue
		}

		// check signature + state of the card
		if card.Txt != "" {
			msg := fmt.Sprintf("register card %s at %d", card.UID, card.Date)
			ck, err := fireblocklib.ECDSAVerify(pkey.Pubkey, msg, card.Signature)
			if err != nil || !ck {
				continue
			}
			// check card
			_, err3 := fireblocklib.VerifyCard(card.Txt, pkey.KeyUID, pkey.KType)
			if err3 != nil {
				continue
			}
		}

		// key state
		if (key.State & 7) != 3 {
			continue
		}
		// check certificate
		message := fmt.Sprintf("%s||%s", hash, key.KeyUID)
		ck, err := fireblocklib.VerifySignature(key.KType, key.Pubkey, message, certificate.Signature)
		if err != nil {
			continue
		}
		if !ck {
			continue
		}
		// check delegation
		message2 := fmt.Sprintf("approved key is %s at %d", key.KeyUID, key.Date)
		ck2, err2 := fireblocklib.VerifySignature("ecdsa", pkey.Pubkey, message2, key.Signature)
		if err2 != nil {
			continue
		}
		if !ck2 {
			continue
		}
		// check metadataSignature
		if certificate.MetadataSignature != "" {
			metadataSID := fireblocklib.Keccak256(certificate.Metadata)
			message3 := fmt.Sprintf("%s||%s||%s", metadataSID, hash, key.KeyUID)
			ck3, err3 := fireblocklib.VerifySignature(key.KType, key.Pubkey, message3, certificate.MetadataSignature)
			if err3 != nil {
				continue
			}
			if !ck3 {
				continue
			}
		}
		validity = true
		break
	}
	if validity {
		verifyExist(pkey, card, filename, hash, verbose)
		os.Exit(0)
	} else {
		verifyError(pkey, card, fireblocklib.InvalidFile, fmt.Sprintf("Not a valid file"), verbose)
		os.Exit(1)
	}
}
