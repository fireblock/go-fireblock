package common

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
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

// ECDSASign sign with ECDSA algo
func ECDSASign(privkey, message string) (string, error) {
	// load key
	priv, err := loadECDSAPrivateKey(privkey)
	if err != nil {
		return "", err
	}
	// compute hash
	hashed := RawSha256(message)
	r, s, err := ecdsa.Sign(rand.Reader, priv, []byte(hashed))
	if err != nil {
		e := NewFBKError("cannot sign", InvalidEncoding)
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

// ECDSAVerify verify a signed message
func ECDSAVerify(pubkey, message, signature string) (bool, error) {
	// load key
	pub, err := loadECDSAPublicKey(pubkey)
	if err != nil {
		return false, err
	}
	hashed := RawSha256(message)
	sig, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		msg := fmt.Sprintf(`base64.RawURLEncoding.DecodeString error %s`, signature)
		e := NewFBKError(msg, InvalidEncoding)
		return false, e
	}
	if len(sig) != 64 {
		e := NewFBKError(`base64.RawURLEncoding.DecodeString length must be 64`, InvalidEncoding)
		return false, e
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])
	res := ecdsa.Verify(pub, hashed, r, s)
	return res, nil
}

// ECDSAFingerprint compute a fingerprint
func ECDSAFingerprint(key string) (string, error) {
	k, err := loadECDSAPublicKey(key)
	if err != nil {
		return "", err
	}
	const ecdsaFingerprintTemplate = `{"crv":"P-256","kty":"EC","x":"%s","y":"%s"}`
	xx := base64.RawURLEncoding.EncodeToString(k.X.Bytes())
	yy := base64.RawURLEncoding.EncodeToString(k.Y.Bytes())
	msg := fmt.Sprintf(ecdsaFingerprintTemplate, xx, yy)
	return Sha1(msg), nil
}

func fromHex(s string) (*big.Int, error) {
	l, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		msg := fmt.Sprintf(`base64.RawURLEncoding.DecodeString error %s`, s)
		e := NewFBKError(msg, InvalidEncoding)
		return nil, e
	}
	r := new(big.Int).SetBytes(l)
	return r, nil
}

func loadECDSAPrivateKey(key string) (*ecdsa.PrivateKey, error) {
	byt := []byte(key)

	var jwk JWKKey
	if err := json.Unmarshal(byt, &jwk); err != nil {
		e := NewFBKError(err.Error(), InvalidJSON)
		return nil, e
	}
	if jwk.Crv != "P-256" {
		msg := fmt.Sprintf(`key format not supported %s`, jwk.Crv)
		e := NewFBKError(msg, InvalidKey)
		return nil, e
	}
	// private key
	k := new(ecdsa.PrivateKey)
	k.PublicKey.Curve = elliptic.P256()
	x, err1 := fromHex(jwk.X)
	y, err2 := fromHex(jwk.Y)
	d, err3 := fromHex(jwk.D)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	if err3 != nil {
		return nil, err3
	}
	k.PublicKey.X = x
	k.PublicKey.Y = y
	k.D = d
	return k, nil
}

func loadECDSAPublicKey(key string) (*ecdsa.PublicKey, error) {
	byt := []byte(key)

	var jwk JWKKey
	if err := json.Unmarshal(byt, &jwk); err != nil {
		e := NewFBKError(err.Error(), InvalidJSON)
		return nil, e
	}
	if jwk.Crv != "P-256" {
		msg := fmt.Sprintf(`key format not supported %s`, jwk.Crv)
		e := NewFBKError(msg, InvalidKey)
		return nil, e
	}
	// public key
	k := new(ecdsa.PublicKey)
	k.Curve = elliptic.P256()
	x, err1 := fromHex(jwk.X)
	y, err2 := fromHex(jwk.Y)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	k.X = x
	k.Y = y
	return k, nil
}

// ECDSAImport import a base64 encoded url key
func ECDSAImport(b64u string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(b64u)
	if err != nil {
		return "", NewFBKError("cannot decode ECDSA key (b64u)", InvalidKey)
	}
	// decode
	var jwk JWKKey
	if err := json.Unmarshal(data, &jwk); err != nil {
		return "", NewFBKError("cannot decode ECDSA Key (jwk)", InvalidKey)
	}
	jwk.Ext = true
	jwk.KeyOps = []string{"sign"}
	// encode
	res, _ := json.Marshal(jwk)
	return string(res), nil
}

// ECDSAExport export a JSON key into a base64url key
func ECDSAExport(privkey string) (string, error) {
	// decode
	var jwk JWKKey
	if err := json.Unmarshal([]byte(privkey), &jwk); err != nil {
		return "", NewFBKError("cannot decode ECDSA Key (jwk)", InvalidKey)
	}
	// Field order is important.
	str := fmt.Sprintf(`{"crv":"%s","d":"%s", "kty":"EC","x":"%s","y":"%s"}`,
		jwk.Crv,
		jwk.D,
		jwk.X,
		jwk.Y,
	)
	return base64.RawURLEncoding.EncodeToString([]byte(str)), nil
}
