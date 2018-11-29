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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

// Metadata meta data
type Metadata struct {
	Kind     string `json:"kind,omitempty"`
	Filename string `json:"filename,omitempty"`
	Type     string `json:"type,omitempty"`
	Size     int64  `json:"size,omitempty"`
}

// MetadataFile extract Metadata from file
func MetadataFile(filepath string) (Metadata, error) {
	metadata := Metadata{"", "", "", 0}
	file, err := os.Open(filepath)
	if err != nil {
		msg := fmt.Sprintf(`No file at %s`, filepath)
		return metadata, NewFBKError(msg, InvalidFile)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		msg := fmt.Sprintf(`No stat info at %s`, filepath)
		return metadata, NewFBKError(msg, InvalidFile)
	}
	// json -> string
	metadata.Filename = path.Base(filepath)
	metadata.Size = fi.Size()
	return metadata, nil
}

const regexPGPPrivKey = `(?s).*(-----BEGIN PGP PRIVATE KEY BLOCK-----.*-----END PGP PRIVATE KEY BLOCK-----).*`
const regexPGPPubKey = `(?s).*(-----BEGIN PGP PUBLIC KEY BLOCK-----.*-----END PGP PUBLIC KEY BLOCK-----).*`

// ECDSAJWK ecdsa jwk
type ECDSAJWK struct {
	Crv    string   `json:"crv"`
	Kty    string   `json:"kty"`
	X      string   `json:"x"`
	Y      string   `json:"y"`
	D      string   `json:"d,omitempty"`
	Keyops []string `json:"key_ops,omitempty"`
}

// ECDSAFIO ecdsa fio
type ECDSAFIO struct {
	Keys []json.RawMessage `json:"keys"`
}

// KeyFIO key fio
type KeyFIO struct {
	Pubkey  string `json:"pubkey"`
	Privkey string `json:"privkey"`
	KeyUID  string `json:"keyuid"`
	KType   string `json:"ktype"`
}

// LoadFioContent load a fio content
func LoadFioContent(content string) (ktype, keyuid, privkey, pubkey string, err error) {
	keys := make([]KeyFIO, 0)
	err = json.Unmarshal([]byte(content), &keys)
	if err != nil {
		j, _ := json.Marshal(&keys)
		fmt.Print(string(j))
		return "", "", "", "", NewFBKError("invalid format file", InvalidKey)
	}
	// try to load the first one
	for _, key := range keys {
		if key.KType != "pgp" && key.KType != "ecdsa" {
			return "", "", "", "", NewFBKError("invalid ktype", InvalidKey)
		}
		return key.KType, key.KeyUID, key.Privkey, key.Pubkey, nil
	}
	return "", "", "", "", NewFBKError("no key in fio file", InvalidKey)
}

// LoadFioFile load a fio file
func LoadFioFile(filepath string) (string, string, string, string, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		msg := fmt.Sprintf(`No file at %s`, filepath)
		return "", "", "", "", NewFBKError(msg, InvalidFile)
	}
	return LoadFioContent(string(content))
}

// LoadB64U load a base64url private key
func LoadB64U(data string) (string, string, error) {
	content, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return "", "", NewFBKError("can't decode key", InvalidKey)
	}
	regexPriv, _ := regexp.Compile(regexPGPPrivKey)
	priv := regexPriv.FindStringSubmatch(string(content))
	if len(priv) > 0 {
		fp, err := PGPFingerprint(priv[1])
		if err != nil {
			return "", "", NewFBKError("Malformed private pgp key", InvalidKey)
		}
		keyuid := PGPToB32(fp)
		return keyuid, priv[1], nil
	}
	var k ECDSAJWK
	err = json.Unmarshal(content, &k)
	if err != nil {
		return "", "", NewFBKError("unknown format", InvalidKey)
	}
	if k.D == "" {
		return "", "", NewFBKError("Not a private ECDSA key", InvalidKey)
	}
	fp, err := ECDSAFingerprint(string(content))
	if err != nil {
		return "", "", NewFBKError("no fingerprint", InvalidKey)
	}
	return fp, string(content), nil
}

// VerifySignature verify signature
func VerifySignature(ktype, pubkey, message, signature string) (bool, error) {
	if ktype == "ecdsa" {
		return ECDSAVerify(pubkey, message, signature)
	} else if ktype == "pgp" {
		return PGPVerify(pubkey, message, signature)
	}
	msg := fmt.Sprintf("Invalid key type %s", ktype)
	return false, NewFBKError(msg, InvalidKey)
}

// ListFilesInDirectory recursive
func ListFilesInDirectory(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
