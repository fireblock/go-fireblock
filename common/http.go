package common

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ServerURL is the url of the service
var ServerURL = "https://fireblock.io"

// JSONRes is a common struct used in json response
type JSONRes struct {
	Errors []ErrorRes      `json:"errors"`
	Data   json.RawMessage `json:"data"`
}

// ErrorRes is a common struct used in json errors
type ErrorRes struct {
	ID     string `json:"id"`
	Detail string `json:"detail"`
}

// KeyResponseData is a common struct used in success response
type KeyResponseData struct {
	ID  string        `json:"id"`
	Key []interface{} `json:"key"`
}

// HTTPKey API request on fireblock.io for a key
func HTTPKey(keyuid string) (string, error) {
	res, _ := http.Get(ServerURL + "/api/key?keyuid=" + keyuid)
	var response JSONRes
	json.NewDecoder(res.Body).Decode(&response)
	// check errors in response
	if len(response.Errors) > 0 {
		msg := fmt.Sprintf(`No key %s found`, keyuid)
		return "", NewFBKError(msg, InvalidKey)
	}
	var data KeyResponseData
	json.Unmarshal(response.Data, &data)
	if len(data.Key) != 4 {
		msg := fmt.Sprintf(`Invalid key %s`, keyuid)
		return "", NewFBKError(msg, InvalidKey)
	}
	pub := data.Key[0].(string)
	// no need to check fp data.Key[1].(string)
	revoked := data.Key[2].(bool)
	// no need to check if key is closed
	if revoked {
		msg := fmt.Sprintf(`Key %s revoked`, keyuid)
		return "", NewFBKError(msg, InvalidKey)
	}
	return pub, nil
}

// CardReq http request
type CardReq struct {
	Keyuid string `json:"keyuid"`
}

// CardData http request struct
type CardData struct {
	ID   string `json:"id"`
	Card string `json:"card"`
}

// HTTPCard API
func HTTPCard(keyuid, token string) (string, error) {
	req := CardReq{keyuid}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	res, _ := postWithToken(http.DefaultClient, ServerURL+"/api/internal/card", token, buffer)
	if res.StatusCode == 403 {
		return "", NewFBKError("Invalid token provided", InvalidCard)
	}
	var response JSONRes
	json.NewDecoder(res.Body).Decode(&response)
	// check errors in response
	if len(response.Errors) > 0 {
		err := response.Errors[0]
		return "", NewFBKError(err.ID, InvalidCard)
	}
	var data CardData
	json.Unmarshal(response.Data, &data)
	if data.ID == "success" {
		return data.Card, nil
	}
	return "", NewFBKError("", InvalidCard)
}

// SignReq http request struct
type SignReq struct {
	Hash      string `json:"hash"`
	Keyuid    string `json:"keyuid"`
	Signature string `json:"signature"`
	Metadata  string `json:"metadata"`
}

// SignData http request struct
type SignData struct {
	ID    string `json:"id"`
	Proof string `json:"proof"`
}

// HTTPSign sign
func HTTPSign(token, hash, keyuid, signature, metadata string) (string, error) {
	sig64 := base64.StdEncoding.EncodeToString([]byte(signature))
	req := SignReq{hash, keyuid, sig64, metadata}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req)
	res, _ := postWithToken(http.DefaultClient, ServerURL+"/api/create-proof", token, buffer)
	if res.StatusCode == 403 {
		return "", NewFBKError("Invalid token provided", InvalidSignature)
	}
	var response JSONRes
	json.NewDecoder(res.Body).Decode(&response)
	// check errors in response
	if len(response.Errors) > 0 {
		err := response.Errors[0]
		return "", NewFBKError(err.ID, InvalidSignature)
	}
	var data SignData
	json.Unmarshal(response.Data, &data)
	if data.ID == "success" {
		return data.Proof, nil
	}
	return "", NewFBKError("", InvalidSignature)
}

func postWithToken(c *http.Client, url string, token string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-access-token", token)
	return c.Do(req)
}
