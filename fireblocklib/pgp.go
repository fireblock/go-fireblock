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
	"bytes"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/clearsign"
	"golang.org/x/crypto/openpgp/packet"
)

// PGPSign clearsign a message
func PGPSign(message, privkey, passphrase string) (string, error) {
	// load private key
	entity, err := loadPrivateKey(privkey, passphrase)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	plaintext, err := clearsign.Encode(&buf, entity.PrivateKey, nil)
	if err != nil {
		return "", NewFBKError("cannot create a signature", InvalidSignature)
	}
	plaintext.Write([]byte(message))
	plaintext.Close()
	return buf.String(), nil
}

// PGPVerify a clearsign message
func PGPVerify(signature, message, pubkey string) (bool, error) {
	entity, err := loadPGPPublicKey(pubkey)
	if err != nil {
		return false, err
	}
	// load clearsigned message and extract signature
	bck, remain := clearsign.Decode([]byte(signature))
	if len(remain) != 0 {
		msg := fmt.Sprintf("Not the signature attended: %s", signature)
		return false, NewFBKError(msg, InvalidSignature)
	}

	if bck.ArmoredSignature == nil {
		return false, NewFBKError("No signature found", InvalidSignature)
	}

	block := bck.ArmoredSignature
	if block.Type != openpgp.SignatureType {
		return false, NewFBKError("No armored part in signature", InvalidSignature)
	}

	reader := packet.NewReader(block.Body)
	pkt, err := reader.Next()
	if err != nil {
		return false, NewFBKError("Cannot read armored part", InvalidSignature)
	}
	sig, ok := pkt.(*packet.Signature)
	if !ok {
		return false, NewFBKError("Cannot read armored part", InvalidSignature)
	}
	hash := sig.Hash.New()
	_, err = io.Copy(hash, bytes.NewBufferString(message))
	if err != nil {
		return false, NewFBKError("Cannot compute hash in armored part", InvalidSignature)
	}
	err = entity.PrimaryKey.VerifySignature(hash, sig)
	if err != nil {
		return false, NewFBKError("Signature doesn't match", InvalidSignature)
	}
	return true, nil
}

// PGPFingerprint get fingerprint of pgp key
func PGPFingerprint(key string) (string, error) {
	entitylist, err := openpgp.ReadArmoredKeyRing(bytes.NewBufferString(key))
	if err != nil {
		return "", NewFBKError(err.Error(), InvalidKey)
	}
	// use only the first key
	entity := entitylist[0]
	fp := fmt.Sprintf("0x%x", entity.PrimaryKey.Fingerprint)
	return fp, nil
}

// PGPExport pgp armored priv key into base64url string
func PGPExport(key string) string {
	/*	entity, err := loadPrivateKey(key, passphrase)
		if err != nil {
			return "", err
		}
		buf := new(bytes.Buffer)
		entity.SerializePrivate(buf, nil)
		res := base64.RawURLEncoding.EncodeToString(buf.Bytes())
		return res, nil
	*/
	return base64.RawURLEncoding.EncodeToString([]byte(key))
}

// PGPImport toto
func PGPImport(key string) (string, error) {
	/*
		el, err := openpgp.ReadKeyRing(strings.NewReader(string(serialized)))
		if err != nil {
			return "", err
		}
		entity := el[0]
		// on encode la clef
		buf := new(bytes.Buffer)
		w, err := armor.Encode(buf, openpgp.PrivateKeyType, nil)
		if err != nil {
			return "", err
		}
		defer w.Close()
		err = entity.SerializePrivate(w, nil)
		if err != nil {
			return "", err
		}
		return buf.String(), nil
	*/
	serialized, err := base64.RawURLEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	return string(serialized), nil
}

func loadPGPPublicKey(pubkey string) (*openpgp.Entity, error) {
	entitylist, err := openpgp.ReadArmoredKeyRing(bytes.NewBufferString(pubkey))
	if err != nil {
		return nil, NewFBKError(err.Error(), InvalidKey)
	}
	// use only the first key
	entity := entitylist[0]
	return entity, nil
}

func loadPrivateKey(privkey, passphrase string) (*openpgp.Entity, error) {
	entitylist, err := openpgp.ReadArmoredKeyRing(bytes.NewBufferString(privkey))
	if err != nil {
		return nil, NewFBKError(err.Error(), InvalidKey)
	}
	// use only the first key
	entity := entitylist[0]
	// check if the private key is encrypted
	if entity.PrivateKey != nil && entity.PrivateKey.Encrypted {
		err := entity.PrivateKey.Decrypt([]byte(passphrase))
		if err != nil {
			return nil, NewFBKError(err.Error(), InvalidPassphrase)
		}
	}
	// decrypt subkeys
	for _, subkey := range entity.Subkeys {
		if subkey.PrivateKey != nil && subkey.PrivateKey.Encrypted {
			err := subkey.PrivateKey.Decrypt([]byte(passphrase))
			if err != nil {
				return nil, NewFBKError(err.Error(), InvalidPassphrase)
			}
		}
	}
	return entity, nil
}
