package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fireblock/go-fireblock/common"
)

var SERVER_URL string

// ProjectVerifyResponse http request
type HttpResponse struct {
	Errors []common.ErrorRes `json:"errors,omitempty"`
	Data   json.RawMessage   `json:"data"`
}

func SetServerURL(url string) {
	SERVER_URL = url
}

func CreateURL(uri string) string {
	url := SERVER_URL + uri
	return url
}

func Post(url string, param interface{}) (json.RawMessage, error) {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(param)
	res, err := http.Post(url, "application/json; charset=utf-8", buffer)
	if err != nil {
		return nil, common.NewFBKError("url: "+url, common.NetworkError)
	}
	var response HttpResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, common.NewFBKError(fmt.Sprintf("http response error %s", url), common.APIError)
	}
	// check result
	if len(response.Errors) > 0 {
		message := fmt.Sprintf("Project Error: %s %s", response.Errors[0].ID, response.Errors[0].Detail)
		return nil, common.NewFBKError(message, common.APIError)
	}
	return response.Data, nil
}

func Get() {

}
