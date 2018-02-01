package common

import (
	"bytes"
	"errors"
	"io"
	"regexp"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

var sigWithDataReg = regexp.MustCompile(`-----BEGIN.*\nHash:.*\n\n(.*)\n(-----BEGIN[\s\S]+)`)

// Verify a pgp signature
func Verify(signatureIncludingData []byte, pubkey [][]byte) (bool, error) {
	regexRes := sigWithDataReg.FindAllSubmatch(signatureIncludingData, 1)
	if len(regexRes) == 1 && len(regexRes[0]) == 3 {
		return VerifyDetached(regexRes[0][1], regexRes[0][2], pubkey)
	}
	return false, errors.New("couldn't parse the signed data of your input")
}

// VerifyDetached a pgp signature
func VerifyDetached(data, signature []byte, pubKey [][]byte) (bool, error) {
	return VerifyStream(bytes.NewBuffer(data), bytes.NewBuffer(signature), pubKey)
}

// VerifyStream verify a pgp stream
func VerifyStream(dataReader, signatureReader io.Reader, pubKey [][]byte) (bool, error) {
	entitylist, err := readPGPKeys(pubKey)
	if err != nil {
		return false, err
	}

	block, err := armor.Decode(signatureReader)

	if block.Type != openpgp.SignatureType {
		return false, errors.New("Invalid signature")
	}

	reader := packet.NewReader(block.Body)
	pkt, err := reader.Next()
	if err != nil {
		return false, err
	}

	sig, ok := pkt.(*packet.Signature)
	if !ok {
		return false, errors.New("Invalid signature")
	}
	hash := sig.Hash.New()
	_, err = io.Copy(hash, dataReader)
	if err != nil {
		return false, err
	}
	err = entitylist[0].PrimaryKey.VerifySignature(hash, sig)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ReadKeys read a pgp key (public or private)
func readPGPKeys(keys [][]byte) (el openpgp.EntityList, err error) {
	for _, key := range keys {
		entitylist, err := openpgp.ReadArmoredKeyRing(bytes.NewBuffer(key))
		if err != nil {
			return nil, err
		}
		for _, e := range entitylist {
			el = append(el, e)
		}
	}
	return el, nil
}
