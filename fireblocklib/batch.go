// Copyright 2018 The go-fireblock Authors
// This file is part of the go-fireblock library.
//
// The go-fireblock library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-fireblock library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-fireblock library. If not, see <http://www.gnu.org/licenses/>.
//
// Package common contains various helper functions.

package fireblocklib

import "encoding/json"

// BatchElem batch element
type BatchElem struct {
	Hash     string `json:"h"`
	Filename string `json:"f,omitempty"`
	Type     string `json:"t,omitempty"`
	Size     int64  `json:"s"`
}

func readBatch(data string) ([]BatchElem, error) {
	byt := []byte(data)
	var result []BatchElem
	if err := json.Unmarshal(byt, &result); err != nil {
		e := NewFBKError(err.Error(), InvalidJSON)
		return result, e
	}
	return result, nil
}

func isIn(batch []BatchElem, hash string) bool {
	for _, element := range batch {
		if element.Hash == hash {
			return true
		}
	}
	return false
}

// IsInBatch check if hash is in batch
func IsInBatch(data, hash string) bool {
	dataJ, err := readBatch(data)
	if err != nil {
		return false
	}
	return isIn(dataJ, hash)
}
