package common

import (
	"bytes"
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
	entity, err := loadPublicKey(pubkey)
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

func loadPublicKey(pubkey string) (*openpgp.Entity, error) {
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
