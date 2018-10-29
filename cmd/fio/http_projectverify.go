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

func fbkError(err error, verbose bool) {
	e := err.(*fireblocklib.FBKError)
	if e != nil {
		fmt.Printf("code: %d detail: %s\n", e.Type(), e.Error())
		os.Exit(1)
	}
	fmt.Printf(err.Error())
	os.Exit(1)
}

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

func projectVerify(server, filename, hash, pkeyUID string, verbose bool) {
	// create url request
	SetServerURL(server)
	url := CreateURL("/api/verify-by-project")

	// json inputs + request
	req := ProjectVerifyReq{hash, pkeyUID}
	res, err := Post(url, req)
	if err != nil {
		fbkError(err, verbose)
		os.Exit(1)
	}

	// parse output
	var response ProjectVerifyValueReturn
	err = json.Unmarshal(res, &response)
	if err != nil {
		j, _ := json.Marshal(&res)
		fmt.Print(string(j))
		os.Exit(1)
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

		certificateHash := hash
		if certificate.Hash != hash {
			certificateHash = certificate.Hash
			batch := certificate.Batch
			hash2 := fireblocklib.Sha256(batch)
			if hash2 != certificate.Hash {
				continue
			}
			batchArray, err2 := fireblocklib.ReadBatch(certificate.Batch)
			if err2 != nil {
				continue
			}
			found := false
			for _, element := range batchArray {
				if element.Hash == hash {
					found = true
				}
			}
			if !found {
				continue
			}
		}

		// check signature + state of the card
		if card.Txt != "" {
			msg := fmt.Sprintf("register card %s at %d", card.UID, card.Date)
			ck, err := fireblocklib.ECDSAVerify(pkey.Pubkey, msg, card.Signature)
			if err != nil || !ck {
				fbkError(fireblocklib.NewFBKError(fmt.Sprintf("Project Error: invalid signature of the card"), fireblocklib.InvalidProject), verbose)
			}
			// check card
			_, err3 := fireblocklib.VerifyCard(card.Txt, pkey.KeyUID, pkey.KType)
			if err3 != nil {
				fbkError(err3, verbose)
			}
		}

		// check pkey state
		if (pkey.State & 15) != 3 {
			fbkError(fireblocklib.NewFBKError(fmt.Sprintf("Project Error: invalid pkey state"), fireblocklib.InvalidProject), verbose)
		}

		// check key state
		if (key.State & 7) != 3 {
			continue
		}
		// check certificate
		message := fmt.Sprintf("%s||%s", certificateHash, key.KeyUID)
		ck, err := fireblocklib.VerifySignature(key.KType, key.Pubkey, message, certificate.Signature)
		if err != nil {
			continue
		}
		if !ck {
			continue
		}
		// check delegation
		message2 := fmt.Sprintf("approved key is %s at %d", key.KeyUID, key.Date)
		ck2, err2 := fireblocklib.ECDSAVerify(pkey.Pubkey, message2, key.Signature)
		if err2 != nil {
			continue
		}
		if !ck2 {
			continue
		}
		// check metadataSignature
		if certificate.MetadataSignature != "" {
			metadataSID := fireblocklib.Keccak256(certificate.Metadata)
			message3 := fmt.Sprintf("%s||%s||%s", metadataSID, certificateHash, key.KeyUID)
			ck3, err3 := fireblocklib.VerifySignature(key.KType, key.Pubkey, message3, certificate.MetadataSignature)
			if err3 != nil {
				continue
			}
			if !ck3 {
				continue
			}
		}
		// all verification completed
		validity = true
		break
	}
	if validity {
		verifySuccess(pkey, card, filename, hash, verbose)
		os.Exit(0)
	} else {
		verifyError(pkey, card, fireblocklib.InvalidFile, fmt.Sprintf("Not a valid file"), verbose)
		os.Exit(1)
	}
}
