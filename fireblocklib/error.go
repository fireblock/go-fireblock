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

// Errors
const (
	None        = 100
	InvalidType = 200

	InvalidAuthentication = 300
	InvalidKey            = 301
	InvalidCard           = 302
	InvalidProof          = 303
	InvalidSignature      = 304
	InvalidPassphrase     = 305
	InvalidProject        = 306
	InvalidFile           = 310
	InvalidJSON           = 320
	InvalidEncoding       = 321
	APIError              = 350

	NetworkError    = 400
	NetworkError404 = 404

	AlreadyExist = 500

	UnknownError    = 600
	UnknownElement  = 601
	UnknownProvider = 602

	EthOpError      = 800
	EthNotEnoughGas = 801

	NotYetImplemented = 999
)

// FBKError struct
type FBKError struct {
	err  string
	code int
}

// NewFBKError default Constructor
func NewFBKError(text string, code int) *FBKError {
	return &FBKError{text, code}
}

func (e *FBKError) Error() string {
	return e.err
}

// Type type
func (e *FBKError) Type() int {
	return e.code
}
