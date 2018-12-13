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

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fireblockPubKey string
var fireblockPrivKey string
var fireblockIDPub string
var fireblockMessage string
var fireblockSignature string

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	bpub, _ := ioutil.ReadFile("testdata/pub_99090EAE43316B2BA65EC52BCD5834A3E07EDB2C.pgp.asc")
	fireblockPubKey = string(bpub)
	bpriv, _ := ioutil.ReadFile("testdata/priv_99090EAE43316B2BA65EC52BCD5834A3E07EDB2C.pgp.asc")
	fireblockPrivKey = string(bpriv)
	bpub2, _ := ioutil.ReadFile("testdata/pub_7B173475CB00C65509513D98C8BC988E1F21EB77.pgp.asc")
	fireblockIDPub = string(bpub2)
	bmsg, _ := ioutil.ReadFile("testdata/message.txt")
	fireblockMessage = string(bmsg)
	bsig, _ := ioutil.ReadFile("testdata/message.txt.asc")
	fireblockSignature = string(bsig)
}

func shutdown() {}

func TestDetachedSign(t *testing.T) {
	msg := "un message a signer doit faire une longueur"
	sig, err := PGPSign(fireblockPrivKey, msg, "fireblock")
	assert.Equal(t, err, nil, "no error when signing")
	res, _ := PGPVerify(fireblockPubKey, msg, sig)
	assert.Equal(t, res, true, "verify matched")
}

func TestDetachedSign2(t *testing.T) {
	_, err := PGPSign("a bad key", "un message", "fireblock")
	assert.NotNil(t, err, "bad private key")
}

func TestDetachedSign3(t *testing.T) {
	res, _ := PGPVerify(fireblockPubKey, fireblockMessage, fireblockSignature)
	assert.Equal(t, res, true, "verify matched")
}

func TestPGPGetFingerprint(t *testing.T) {
	fp, _ := PGPFingerprint(fireblockPrivKey)
	assert.Equal(t, fp, "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c", "Invalid fingerprint")
	fp, _ = PGPFingerprint(fireblockPubKey)
	assert.Equal(t, fp, "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c", "Invalid fingerprint")
	_, err := PGPFingerprint("not a pgp key")
	e := err.(*FBKError)
	assert.Equal(t, e.Type(), InvalidKey, "Invalid fingerprint")
}
func TestPGPPrivKeyToB64U1(t *testing.T) {
	res := PGPExport(fireblockPrivKey)
	assert.NotNil(t, res, "not empty string")
	res2, _ := PGPImport(res)
	assert.NotNil(t, res2, "not empty string")
}
