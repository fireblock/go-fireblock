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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fireblock/go-fireblock/fireblocklib"
)

// ServerURL server url
var ServerURL string

// HTTPResponse http request
type HTTPResponse struct {
	Errors []fireblocklib.ErrorRes `json:"errors,omitempty"`
	Data   json.RawMessage         `json:"data"`
}

// SetServerURL set server url
func SetServerURL(url string) {
	ServerURL = url
}

// CreateURL create an uri
func CreateURL(uri string) string {
	url := ServerURL + uri
	return url
}

// Post request
func Post(url string, param interface{}) (json.RawMessage, error) {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(param)
	res, err := http.Post(url, "application/json; charset=utf-8", buffer)
	if err != nil {
		return nil, fireblocklib.NewFBKError("url: "+url, fireblocklib.NetworkError)
	}
	// if our API return 200, we have a json result
	if res.StatusCode == 200 {
		var response json.RawMessage
		err = json.NewDecoder(res.Body).Decode(&response)
		return response, nil
	}
	// if error, check if we have a json error
	var response HTTPResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fireblocklib.NewFBKError(fmt.Sprintf("http response error %s", url), fireblocklib.APIError)
	}
	// check result
	if len(response.Errors) > 0 {
		message := fmt.Sprintf("Error %d: %s\n", res.StatusCode, response.Errors[0].ID)
		message += fmt.Sprintf("Detail: %s", response.Errors[0].Detail)
		return nil, fireblocklib.NewFBKError(message, fireblocklib.APIError)
	}
	return response.Data, nil
}

func Get() {

}
