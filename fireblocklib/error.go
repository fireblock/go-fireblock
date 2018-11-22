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
	None = 0

	InvalidAuthentication = 10
	InvalidKey            = 11
	InvalidCard           = 12
	InvalidProof          = 13
	InvalidSignature      = 14
	InvalidPassphrase     = 15
	InvalidProject        = 16
	InvalidHash           = 17
	InvalidFile           = 18
	InvalidJSON           = 19
	InvalidEncoding       = 21
	NetworkError          = 30
	NetworkError404       = 31
	NoFile                = 32
	SignError             = 40
	APIError              = 50
	AlreadyExist          = 60

	UnknownError    = 90
	UnknownElement  = 91
	UnknownProvider = 92

	EthOpError      = 100
	EthNotEnoughGas = 101

	NotYetImplemented = 127
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
