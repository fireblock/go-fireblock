package common

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/fireblock/go-fireblock/common/errors"
)

// JWKKey a key
type JWKKey struct {
	D      string   `json:"d,omitempty"`
	Crv    string   `json:"crv"`
	Ext    bool     `json:"ext"`
	KeyOps []string `json:"key_ops"`
	Kty    string   `json:"kty"`
	X      string   `json:"x"`
	Y      string   `json:"y"`
}

// ECDSAReadKeys read ECDSA Keys
func ECDSAReadKeys(key string) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	byt := []byte(key)

	var jwk JWKKey
	if err := json.Unmarshal(byt, &jwk); err != nil {
		e := errors.NewFBKError(err.Error(), errors.InvalidJson)
		return nil, nil, e
	}
	if jwk.Crv != "P-256" {
		msg := fmt.Sprintf(`key format not supported %s`, jwk.Crv)
		e := errors.NewFBKError(msg, errors.InvalidKey)
		return nil, nil, e
	}
	if jwk.D == "" {
		// public key
		k := new(ecdsa.PublicKey)
		k.Curve = elliptic.P256()
		x, err1 := fromHex(jwk.X)
		y, err2 := fromHex(jwk.Y)
		if err1 != nil {
			return nil, nil, err1
		}
		if err2 != nil {
			return nil, nil, err2
		}
		k.X = x
		k.Y = y
		return k, nil, nil
	} else {
		// private key
		key := new(ecdsa.PrivateKey)
		key.PublicKey.Curve = elliptic.P256()
		x, err1 := fromHex(jwk.X)
		y, err2 := fromHex(jwk.Y)
		d, err3 := fromHex(jwk.D)
		if err1 != nil {
			return nil, nil, err1
		}
		if err2 != nil {
			return nil, nil, err2
		}
		if err3 != nil {
			return nil, nil, err3
		}
		key.PublicKey.X = x
		key.PublicKey.Y = y
		key.D = d
		return &key.PublicKey, key, nil
	}
}

func ECDSASign(priv *ecdsa.PrivateKey, message string) (string, error) {
	// compute hash
	hashed := RawSha256(message)
	r, s, err := ecdsa.Sign(rand.Reader, priv, []byte(hashed))
	if err != nil {
		e := errors.NewFBKError("cannot sign", errors.InvalidEncoding)
		return "", e
	}
	// create r,s signature into a 64 []byte
	rBytes := r.Bytes()
	rBytesPadded := make([]byte, 32)
	copy(rBytesPadded[32-len(rBytes):], rBytes)
	sBytes := s.Bytes()
	sBytesPadded := make([]byte, 32)
	copy(sBytesPadded[32-len(sBytes):], sBytes)
	out := append(rBytesPadded, sBytesPadded...)
	// encode
	res := base64.RawURLEncoding.EncodeToString(out)
	return res, nil
}

func ECDSAVerify(pub *ecdsa.PublicKey, message, signature string) (bool, error) {
	hashed := RawSha256(message)
	sig, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		msg := fmt.Sprintf(`base64.RawURLEncoding.DecodeString error %s`, signature)
		e := errors.NewFBKError(msg, errors.InvalidEncoding)
		return false, e
	}
	if len(sig) != 64 {
		e := errors.NewFBKError(`base64.RawURLEncoding.DecodeString length must be 64`, errors.InvalidEncoding)
		return false, e
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])
	res := ecdsa.Verify(pub, hashed, r, s)
	return res, nil
}

func ECDSAFingerprint(x, y *big.Int) string {
	const ecdsaFingerprintTemplate = `{"crv":"P-256","kty":"EC","x":"%s","y":"%s"}`
	xx := base64.RawURLEncoding.EncodeToString(x.Bytes())
	yy := base64.RawURLEncoding.EncodeToString(y.Bytes())
	msg := fmt.Sprintf(ecdsaFingerprintTemplate, xx, yy)
	return Sha1(msg)
}

func fromHex(s string) (*big.Int, error) {
	l, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		msg := fmt.Sprintf(`base64.RawURLEncoding.DecodeString error %s`, s)
		e := errors.NewFBKError(msg, errors.InvalidEncoding)
		return nil, e
	}
	r := new(big.Int).SetBytes(l)
	return r, nil
}
