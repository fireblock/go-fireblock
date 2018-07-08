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

package common

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

// Sha256File compute the sha256 of a file
func Sha256File(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		msg := fmt.Sprintf(`No file at %s`, filepath)
		return "", NewFBKError(msg, InvalidFile)
	}
	defer file.Close()
	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", NewFBKError("internal", InvalidFile)
	}
	return "0x" + hex.EncodeToString(hasher.Sum(nil)), nil
}

// Metadata meta data
type Metadata struct {
	Filename string `json:"filename"`
	Type     string `json:"type,omitempty"`
	Size     int64  `json:"size"`
}

// MetadataFile extract Metadata from file
func MetadataFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		msg := fmt.Sprintf(`No file at %s`, filepath)
		return "", NewFBKError(msg, InvalidFile)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		msg := fmt.Sprintf(`No stat info at %s`, filepath)
		return "", NewFBKError(msg, InvalidFile)
	}
	// json -> string
	metadata := Metadata{"", "", 0}
	metadata.Filename = path.Base(filepath)
	metadata.Size = fi.Size()
	export, _ := json.Marshal(metadata)
	return string(export), nil
}

// Keccak256 return keccak256 value of a string
func Keccak256(text string) string {
	if len(text) == 0 {
		return "0x0"
	}
	hasher := sha3.NewKeccak256()
	hasher.Write([]byte(text))
	return "0x" + hex.EncodeToString(hasher.Sum(nil))
}

// Sha256 return sha256 value of a string
func Sha256(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return "0x" + hex.EncodeToString(hasher.Sum(nil))
}

// RawSha256 return sha256 without '0x' prefix
func RawSha256(text string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hasher.Sum(nil)
}

// IsSha256 return true if it's a sha256
func IsSha256(value string) bool {
	if len(value) > 66 {
		return false
	}
	b, _ := regexp.MatchString("^0x[0-9A-Fa-f]{0,128}$", value)
	return b
}

// Sha1 return sha1 value of a string
func Sha1(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	return "0x" + hex.EncodeToString(hasher.Sum(nil))
}

// PGPToB32 convert PGP key to bytes32
func PGPToB32(pgp string) string {
	if strings.Index(pgp, "0x") == 0 {
		r := pgp[2:]
		return "0x100000000000000000000000" + r
	}
	return "0x100000000000000000000000" + pgp
}

// B32ToPGP convert bytes32 to pgp
func B32ToPGP(b32 string) string {
	return "0x" + b32[26:]
}

// ECDSAToB32 convert ecdsa to b32
func ECDSAToB32(ecdsa string) string {
	if strings.Index(ecdsa, "0x") == 0 {
		r := ecdsa[2:]
		return "0x200000000000000000000000" + r
	}
	return "0x200000000000000000000000" + ecdsa
}

// B32ToECDSA convert B32 to ECDSA
func B32ToECDSA(b32 string) string {
	return "0x" + b32[26:]
}

// B32Type return type of key
func B32Type(b32 string) string {
	if len(b32) != 66 {
		return "unknown"
	} else if strings.Index(b32, "0x200000000000000000000000") == 0 {
		return "ecdsa"
	} else if strings.Index(b32, "0x100000000000000000000000") == 0 {
		return "pgp"
	} else if strings.Index(b32, "0x000000000000000000000000") == 0 {
		return "eth"
	} else {
		return "unknown"
	}
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

// LoadFioContent load a fio content
func LoadFioContent(content string) (string, string, string, error) {
	// check if PGP
	regexPriv, _ := regexp.Compile(regexPGPPrivKey)
	regexPub, _ := regexp.Compile(regexPGPPubKey)
	priv := regexPriv.FindStringSubmatch(content)
	pub := regexPub.FindStringSubmatch(content)
	if len(priv) > 0 || (len(pub) > 0) {
		var err error
		fp := ""
		privkey := ""
		pubkey := ""
		if len(priv) > 0 {
			privkey = priv[1]
			fp, err = PGPFingerprint(privkey)
			if err != nil {
				return "", "", "", NewFBKError("no fingerprint", InvalidKey)
			}
		}
		if len(pub) > 0 {
			pubkey = pub[1]
			fp, err = PGPFingerprint(pubkey)
			if err != nil {
				return "", "", "", NewFBKError("no fingerprint", InvalidKey)
			}
		}
		keyuid := PGPToB32(fp)
		return keyuid, privkey, pubkey, nil
	}
	// check if JWK ECDSA
	var data ECDSAFIO
	err := json.Unmarshal([]byte(content), &data)
	if err == nil {
		keyuid := ""
		privkey := ""
		pubkey := ""
		for _, key := range data.Keys {
			txt := string(key)
			var k ECDSAJWK
			err := json.Unmarshal([]byte(txt), &k)
			if err != nil {
				return "", "", "", NewFBKError("unknown format", InvalidKey)
			}
			fp, err := ECDSAFingerprint(txt)
			if err != nil {
				return "", "", "", NewFBKError("no fingerprint", InvalidKey)
			}
			if k.D == "" {
				pubkey = txt
			} else {
				privkey = txt
			}
			keyuid = ECDSAToB32(fp)
		}
		return keyuid, privkey, pubkey, nil
	}
	return "", "", "", NewFBKError("invalid format", InvalidKey)
}

// LoadFioFile load a fio file
func LoadFioFile(filepath string) (string, string, string, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		msg := fmt.Sprintf(`No file at %s`, filepath)
		return "", "", "", NewFBKError(msg, InvalidFile)
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
