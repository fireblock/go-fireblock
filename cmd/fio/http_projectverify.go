package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fireblock/go-fireblock/common"
)

func fbkError(err error, verbose bool) {
	e := err.(*common.FBKError)
	if e != nil {
		fmt.Printf("code: %d detail: %s\n", e.Type(), e.Error())
		os.Exit(1)
	}
	fmt.Printf(err.Error())
	os.Exit(1)
}

func convertPS2CIP(ps common.ProviderState) (cip CardInfoProvider) {
	cip.UID = ps.UID
	cip.Status = ps.Status
	cip.Proof = ps.Proof
	return cip
}

type ProjectVerifyResult struct {
	Key         KeyInfo         `json:"key"`
	Certificate CertificateInfo `json:"certificate"`
}

// ProjectVerifyValue http request
// keyUID, metadata, pkeySignature, certificateSignature, date, cdate, pkeystate: pks, ppubkey, keystate: ks, pubkey, ktype, metadataSignature
type ProjectVerifyValue struct {
	Results []ProjectVerifyResult `json:"results"`
	Card    CardInfo              `json:"card"`
	PKey    KeyInfo               `json:"pkey"`
}

// ProjectVerifyValueReturn http request
type ProjectVerifyValueReturn struct {
	Value ProjectVerifyValue `json:"value"`
	ID    string             `json:"id"`
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

	validity := false
	pkey := response.Value.PKey
	card := response.Value.Card

	// check signature + state of the card
	if card.Txt != "" {
		msg := fmt.Sprintf("register card %s at %d", card.UID, card.Date)
		ck, err := common.ECDSAVerify(pkey.Pubkey, msg, card.Signature)
		if err != nil || !ck {
			fbkError(common.NewFBKError(fmt.Sprintf("Project Error: invalid signature of the card"), common.InvalidProject), verbose)
		}
		// check card
		_, err3 := common.VerifyCard(card.Txt, pkey.KeyUID, pkey.KType)
		if err3 != nil {
			fbkError(err3, verbose)
		}
	}

	// check pkey state
	if (pkey.State & 15) != 3 {
		fbkError(common.NewFBKError(fmt.Sprintf("Project Error: invalid pkey state"), common.InvalidProject), verbose)
	}

	values := response.Value.Results
	for _, value := range values {
		key := value.Key
		certificate := value.Certificate
		// check key state
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
		ck2, err2 := common.ECDSAVerify(pkey.Pubkey, message2, key.Signature)
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
		// all verification completed
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
