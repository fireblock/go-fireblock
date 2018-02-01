package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fireblock/go-fireblock/common/errors"
)

type GlobalVerifyReq struct {
	Hash string `json:"hash"`
	Uuid string `json:"uuid"`
}

type ErrorRes struct {
	Id     string `json:"id"`
	Detail string `json:"detail"`
}

type JsonRes struct {
	Errors []ErrorRes      `json:"errors"`
	Data   json.RawMessage `json:"data"`
}

type GVerifyData struct {
	Id        string          `json:"id"`
	Verified  bool            `json:"verified"`
	CardCheck json.RawMessage `json:"cardCheck"`
	Signature string          `json:"signature"`
	CardId    string          `json:"cardId"`
}

type CardCheckData struct {
	Card       string `json:"card"`
	CardStatus bool   `json:"cardStatus"`
	Key        string `json:"key"`
	Twitter    string `json:"twitter"`
	Github     string `json:"github"`
	Https      string `json:"https"`
	Linkedin   string `json:"linkedin"`
}

func GVerify(hash string, useruid string) {
	req := GlobalVerifyReq{hash, useruid}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	res, _ := http.Post("https://fireblock.io/api/global-verify-proof", "application/json; charset=utf-8", buffer)
	var response JsonRes
	json.NewDecoder(res.Body).Decode(&response)
	// check errors in response
	if len(response.Errors) > 0 {
		err := response.Errors[0]
		fmt.Printf("Error:%s (%s)", err.Id, err.Detail)
		os.Exit(1)
	}
	// no errors
	var data GVerifyData
	json.Unmarshal(response.Data, &data)
	if data.Id == "success" {
		if !data.Verified && data.CardId == "" {
			fmt.Printf("Error: file not registered\n")
			os.Exit(1)
		} else if !data.Verified {
			fmt.Printf("Error: file registered but not valid!\n")
			os.Exit(1)
		} else if data.Verified {
			var cardCheck CardCheckData
			signature := data.Signature
			cardId := data.CardId
			json.Unmarshal(data.CardCheck, &cardCheck)
			if !cardCheck.CardStatus {
				fmt.Printf("Error: file registered but not valid! Card not verified\n")
				os.Exit(1)
			}
			pubkey, ktype, err := CheckAllCard(cardCheck.Card, cardId)
			if err != nil {
				fmt.Printf("Error: file registered but not valid! Card not verified\n")
				os.Exit(1)
			}
			// card verified check the signature
			if len(signature) < 4 || len(signature) > 100000 {
				fmt.Printf("Error: file registered but not valid! Invalid signature\n")
				os.Exit(1)
			}
			if ktype == "pgp" {
				sig := strings.Replace(signature, "\r", "", -1)
				r, err := Verify([]byte(sig), [][]byte{[]byte(pubkey)})
				if err != nil || !r {
					fmt.Printf("Error: file registered but not valid! Invalid signature\n")
					os.Exit(1)
				}
				fmt.Printf("File matched and verified")
				os.Exit(0)
			} else if ktype == "ecdsa" {
				jwkPubKey, _, err2 := ReadECDSAKeys(pubkey)
				if err2 != nil {
					fmt.Printf("Error: file registered but not valid! Invalid ecdsa key\n")
					os.Exit(1)
				}
				message := hash + `||` + cardId
				r, err := ECDSAVerify(jwkPubKey, message, signature)
				if err != nil || !r {
					fmt.Printf("Error: file registered but not valid! Invalid signature\n")
					os.Exit(1)
				}
				fmt.Printf("File matched and verified\n")
				os.Exit(0)
			} else {
				fmt.Printf("Error: file registered but not valid! Signature not supported\n")
				os.Exit(1)
			}
			fmt.Println(err)
			fmt.Println(pubkey)
			fmt.Println(ktype)
			fmt.Printf("%s %s", signature, cardId)
		}
	}
	fmt.Printf("Error: unknown error!\n")
	os.Exit(66)
}

type CardVerifyReq struct {
	Hash   string `json:"hash"`
	CardId string `json:"cardId"`
}

func CVerify(hash string, cardId string) {
	req := CardVerifyReq{hash, cardId}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	res, _ := http.Post("https://fireblock.io/api/verify-proof", "application/json; charset=utf-8", buffer)
	var response JsonRes
	json.NewDecoder(res.Body).Decode(&response)
	// check errors in response
	if len(response.Errors) > 0 {
		err := response.Errors[0]
		fmt.Printf("Error:%s (%s)", err.Id, err.Detail)
		os.Exit(1)
	}
	// no errors
	var data GVerifyData
	json.Unmarshal(response.Data, &data)
	if data.Id == "success" {
		if !data.Verified {
			fmt.Printf("Error: file registered but not valid!\n")
			os.Exit(1)
		} else if data.Verified {
			var cardCheck CardCheckData
			signature := data.Signature
			json.Unmarshal(data.CardCheck, &cardCheck)
			if !cardCheck.CardStatus {
				fmt.Printf("Error: file registered but not valid! Card not verified\n")
				os.Exit(1)
			}
			pubkey, ktype, err := CheckAllCard(cardCheck.Card, cardId)
			if err != nil {
				fmt.Printf("Error: file registered but not valid! Card not verified\n")
				os.Exit(1)
			}
			// card verified check the signature
			if len(signature) < 4 || len(signature) > 100000 {
				fmt.Printf("Error: file registered but not valid! Invalid signature\n")
				os.Exit(1)
			}
			if ktype == "pgp" {
				sig := strings.Replace(signature, "\r", "", -1)
				r, err := Verify([]byte(sig), [][]byte{[]byte(pubkey)})
				if err != nil || !r {
					fmt.Printf("Error: file registered but not valid! Invalid signature\n")
					os.Exit(1)
				}
				fmt.Printf("File matched and verified")
				os.Exit(0)
			} else if ktype == "ecdsa" {
				jwkPubKey, _, err2 := ReadECDSAKeys(pubkey)
				if err2 != nil {
					fmt.Printf("Error: file registered but not valid! Invalid ecdsa key\n")
					os.Exit(1)
				}
				message := hash + `||` + cardId
				r, err := ECDSAVerify(jwkPubKey, message, signature)
				if err != nil || !r {
					fmt.Printf("Error: file registered but not valid! Invalid signature\n")
					os.Exit(1)
				}
				fmt.Printf("File matched and verified\n")
				os.Exit(0)
			} else {
				fmt.Printf("Error: file registered but not valid! Signature not supported\n")
				os.Exit(1)
			}
			fmt.Println(err)
			fmt.Println(pubkey)
			fmt.Println(ktype)
			fmt.Printf("%s %s", signature, cardId)
		}
	}
	fmt.Printf("Error: unknown error!\n")
	os.Exit(66)
}

type KeyResponseData struct {
	Id  string        `json:"id"`
	Key []interface{} `json:"key"`
}

func HTTPKey(keyuid string) (string, error) {
	res, _ := http.Get("https://fireblock.io/api/key?keyuid=" + keyuid)
	var response JsonRes
	json.NewDecoder(res.Body).Decode(&response)
	// check errors in response
	if len(response.Errors) > 0 {
		msg := fmt.Sprintf(`No key %s found`, keyuid)
		return "", errors.NewFBKError(msg, errors.InvalidKey)
	}
	var data KeyResponseData
	json.Unmarshal(response.Data, &data)
	if len(data.Key) != 4 {
		msg := fmt.Sprintf(`Invalid key %s`, keyuid)
		return "", errors.NewFBKError(msg, errors.InvalidKey)
	}
	pub := data.Key[0].(string)
	// no need to check fp data.Key[1].(string)
	revoked := data.Key[2].(bool)
	// no need to check if key is closed
	if revoked {
		msg := fmt.Sprintf(`Key %s revoked`, keyuid)
		return "", errors.NewFBKError(msg, errors.InvalidKey)
	}
	return pub, nil
}
