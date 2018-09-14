package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fireblock/go-fireblock/common"
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
	ID     string                `json:"id"`
	Errors []common.ErrorRes     `json:"errors,omitempty"`
	Data   UserVerifyValueReturn `json:"data"`
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
	sha3uuid := common.Keccak256(useruid)
	req := UserVerifyReq{hash, sha3uuid}
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

		// pkey state
		if (pkey.State & 15) != 3 {
			continue
		}

		// check signature + state of the card
		if card.Txt != "" {
			msg := fmt.Sprintf("register card %s at %d", card.UID, card.Date)
			ck, err := common.ECDSAVerify(pkey.Pubkey, msg, card.Signature)
			if err != nil || !ck {
				continue
			}
			// check card
			_, err3 := common.VerifyCard(card.Txt, pkey.KeyUID, pkey.KType)
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
		ck, err := common.VerifySignature(key.KType, key.Pubkey, message, certificate.Signature)
		if err != nil {
			continue
		}
		if !ck {
			continue
		}
		// check delegation
		message2 := fmt.Sprintf("approved key is %s at %d", key.KeyUID, key.Date)
		ck2, err2 := common.VerifySignature("ecdsa", pkey.Pubkey, message2, key.Signature)
		if err2 != nil {
			continue
		}
		if !ck2 {
			continue
		}
		// check metadataSignature
		if certificate.MetadataSignature != "" {
			metadataSID := common.Keccak256(certificate.Metadata)
			message3 := fmt.Sprintf("%s||%s||%s", metadataSID, hash, key.KeyUID)
			ck3, err3 := common.VerifySignature(key.KType, key.Pubkey, message3, certificate.MetadataSignature)
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
		verifySuccess(pkey, card, filename, hash, verbose)
		os.Exit(0)
	} else {
		verifyError(pkey, card, common.InvalidFile, fmt.Sprintf("Not a valid file"), verbose)
		os.Exit(1)
	}
}
