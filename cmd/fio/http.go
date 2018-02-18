package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/fireblock/go-fireblock/common"
)

// SuccessReturn success data
type SuccessReturn struct {
	Filename string `json:"filename"`
	Verified bool   `json:"verified"`
	Hash     string `json:"hash"`
	UserID   string `json:"user-id"`
	CardID   string `json:"card-id"`
}

// ErrorReturn error return struct
type ErrorReturn struct {
	Error    string `json:"error"`
	Detail   string `json:"detail"`
	Hash     string `json:"hash"`
	Filename string `json:"filename"`
	UserID   string `json:"user-id"`
	CardID   string `json:"card-id"`
}

// GlobalVerifyReq http request struct
type GlobalVerifyReq struct {
	Hash string `json:"hash"`
	UUID string `json:"uuid"`
}

// GVerifyData http request struct
type GVerifyData struct {
	ID        string          `json:"id"`
	Verified  bool            `json:"verified"`
	CardCheck json.RawMessage `json:"cardCheck"`
	Signature string          `json:"signature"`
	CardID    string          `json:"cardId"`
}

// CardCheckData http request
type CardCheckData struct {
	Card       string `json:"card"`
	CardStatus bool   `json:"cardStatus"`
	Key        string `json:"key"`
	Twitter    string `json:"twitter"`
	Github     string `json:"github"`
	HTTPS      string `json:"https"`
	Linkedin   string `json:"linkedin"`
}

// CardVerifyReq http request
type CardVerifyReq struct {
	Hash   string `json:"hash"`
	CardID string `json:"cardId"`
}

func fioSuccess(filename, hash, userid, cardid string, verified, verbose bool) {
	if verbose {
		var r SuccessReturn
		r.Filename = filename
		r.Verified = verified
		r.Hash = hash
		r.UserID = userid
		r.CardID = cardid
		export, _ := json.Marshal(r)
		fmt.Printf("%s\n", export)
	} else {
		fmt.Printf("File matched and verified\n")
	}
	os.Exit(0)
}

func fioError(id, detail, filename, hash, userid, cardid string, verbose bool) {
	if verbose {
		var r ErrorReturn
		r.Error = id
		r.Detail = detail
		r.Hash = hash
		r.UserID = userid
		r.CardID = cardid
		export, _ := json.Marshal(r)
		fmt.Printf("%s\n", export)
	} else {
		fmt.Printf("Error: %s (detail: %s filename: %s hash: %s)\n", id, detail, filename, hash)
	}
	os.Exit(1)
}

// GVerify api global-verify-proof
func GVerify(filename, hash, useruid string, verbose bool) {
	req := GlobalVerifyReq{hash, useruid}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	res, _ := http.Post("https://fireblock.io/api/global-verify-proof", "application/json; charset=utf-8", buffer)
	var response common.JSONRes
	json.NewDecoder(res.Body).Decode(&response)
	// check errors in response
	if len(response.Errors) > 0 {
		err := response.Errors[0]
		fioError(err.ID, err.Detail, filename, hash, useruid, "", verbose)
	}
	// no errors
	var data GVerifyData
	json.Unmarshal(response.Data, &data)
	if data.ID == "success" {
		if !data.Verified && data.CardID == "" {
			fioError("File not registered", "", filename, hash, useruid, "", verbose)
		} else if !data.Verified {
			fioError("File registered but not valid", "", filename, hash, useruid, "", verbose)
		} else if data.Verified {
			var cardCheck CardCheckData
			signature := data.Signature
			cardID := data.CardID
			json.Unmarshal(data.CardCheck, &cardCheck)
			if !cardCheck.CardStatus {
				fioError("File registered but card not verified", "", filename, hash, useruid, cardID, verbose)
			}
			_, pubkey, ktype, err := common.CheckAllCard(cardCheck.Card, cardID)
			if err != nil {
				fioError("File registered but card not verified", "", filename, hash, useruid, cardID, verbose)
			}
			// card verified check the signature
			if len(signature) < 4 || len(signature) > 100000 {
				fioError("File registered but signature not verified", "", filename, hash, useruid, cardID, verbose)
			}
			message := hash + `||` + cardID
			if ktype == "pgp" {
				r, err := common.PGPVerify(signature, message, pubkey)
				if err != nil || !r {
					fioError("File registered but PGP signature not verified", "", filename, hash, useruid, cardID, verbose)
				}
				fioSuccess(filename, hash, useruid, cardID, true, verbose)
			} else if ktype == "ecdsa" {
				jwkPubKey, _, err2 := common.ECDSAReadKeys(pubkey)
				if err2 != nil {
					fioError("File registered but ECDSA signature not verified", "", filename, hash, useruid, cardID, verbose)
				}
				r, err := common.ECDSAVerify(jwkPubKey, message, signature)
				if err != nil || !r {
					fioError("File registered but ECDSA signature not verified", "", filename, hash, useruid, cardID, verbose)
				}
				fioSuccess(filename, hash, useruid, cardID, true, verbose)
			} else {
				fioError("File registered but key format not supported", "", filename, hash, useruid, cardID, verbose)
			}
		}
	}
	fioError("Error unknown", "", filename, hash, useruid, "", verbose)
}

// CVerify api verify-proof
func CVerify(filename, hash, cardID string, verbose bool) {
	req := CardVerifyReq{hash, cardID}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	res, _ := http.Post("https://fireblock.io/api/verify-proof", "application/json; charset=utf-8", buffer)
	var response common.JSONRes
	json.NewDecoder(res.Body).Decode(&response)
	// check errors in response
	if len(response.Errors) > 0 {
		err := response.Errors[0]
		fioError(err.ID, err.Detail, filename, hash, "", cardID, verbose)
	}
	// no errors
	var data GVerifyData
	json.Unmarshal(response.Data, &data)
	if data.ID == "success" {
		if !data.Verified {
			fioError("File registered but not valid", "", filename, hash, "", cardID, verbose)
		} else if data.Verified {
			var cardCheck CardCheckData
			signature := data.Signature
			json.Unmarshal(data.CardCheck, &cardCheck)
			if !cardCheck.CardStatus {
				fioError("File registered but card not verified", "", filename, hash, "", cardID, verbose)
			}
			uuid, pubkey, ktype, err := common.CheckAllCard(cardCheck.Card, cardID)
			useruid := common.Keccak256(uuid)
			if err != nil {
				fioError("File registered but card not verified", "", filename, hash, "", cardID, verbose)
			}
			// card verified check the signature
			if len(signature) < 4 || len(signature) > 100000 {
				fioError("File registered but signature not verified", "", filename, hash, useruid, cardID, verbose)
			}
			message := hash + `||` + cardID
			if ktype == "pgp" {
				r, err := common.PGPVerify(signature, message, pubkey)
				if err != nil || !r {
					fioError("File registered but PGP signature not verified", "", filename, hash, useruid, cardID, verbose)
				}
				fioSuccess(filename, hash, useruid, cardID, true, verbose)
			} else if ktype == "ecdsa" {
				jwkPubKey, _, err2 := common.ECDSAReadKeys(pubkey)
				if err2 != nil {
					fioError("File registered but ECDSA signature not verified", "", filename, hash, useruid, cardID, verbose)
				}
				r, err := common.ECDSAVerify(jwkPubKey, message, signature)
				if err != nil || !r {
					fioError("File registered but ECDSA signature not verified", "", filename, hash, useruid, cardID, verbose)
				}
				fioSuccess(filename, hash, useruid, cardID, true, verbose)
			} else {
				fioError("File registered but key format not supported", "", filename, hash, useruid, cardID, verbose)
			}
		}
	}
	fioError("Error unknown", "", filename, hash, "", cardID, verbose)
}
