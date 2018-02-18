package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	res, _ := http.Get("https://fireblock.io/api/key?keyuid=" + keyuid)
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
