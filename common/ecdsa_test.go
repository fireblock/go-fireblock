package common

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const jwkInvalid = `{ "bad": "json"`

const jwkPubKey1 = `{
	"crv": "P-256",
	"ext": true,
	"key_ops": ["verify"],
	"kty": "EC",
	"x": "qpiv3dvfmLmUEx561WlSWyvMGOFA5r9K8mLt2R7NZzk",
	"y": "FreoygBEeuuxRekKf5g0u-UNhRnfeN5QXGYHEDWpRPQ"
}`

const jwkPrivKey1 = `{
	"crv": "P-256",
	"d": "xr15Ms0pm5VSFLPI76BfUwMZa9J04bUmGlHk1Y0PKKg",
	"ext": true,
	"key_ops": ["sign"],
	"kty": "EC",
	"x": "qpiv3dvfmLmUEx561WlSWyvMGOFA5r9K8mLt2R7NZzk",
	"y": "FreoygBEeuuxRekKf5g0u-UNhRnfeN5QXGYHEDWpRPQ"
}`

const jwkB64U = `eyJjcnYiOiJQLTI1NiIsImQiOiJ4cjE1TXMwcG01VlNGTFBJNzZCZlV3TVphOUowNGJVbUdsSGsxWTBQS0tnIiwgImt0eSI6IkVDIiwieCI6InFwaXYzZHZmbUxtVUV4NTYxV2xTV3l2TUdPRkE1cjlLOG1MdDJSN05aemsiLCJ5IjoiRnJlb3lnQkVldXV4UmVrS2Y1ZzB1LVVOaFJuZmVONVFYR1lIRURXcFJQUSJ9`

const jwkFpKey1 = "0x02e1ee50a71cb8a81aff1461c2d3163b39f88a25"
const jwkKeyuidKey1 = "0x20000000000000000000000002e1ee50a71cb8a81aff1461c2d3163b39f88a25"

const JwkPubKey2 = `{
	"crv": "P-256",
	"ext": true,
	"key_ops": ["verify"],
	"kty": "EC",
	"x": "cD-BiqsSCwm7KfXWM_sRjsl42irUIJybQlk5GlR9ucM",
	"y": "7BKEPCB8gYjrx6p5DuMH0ckLyvn5stprWoUNdwW1L3k"
}`

const JwkPrivKey2 = `{
	"crv": "P-256",
	"d": "nz3bw3pyZDi451Makx2_XzVsFc1jZk7baa-l_5-UhjQ",
	"ext": true,
	"key_ops": ["sign"],
	"kty": "EC",
	"x": "cD-BiqsSCwm7KfXWM_sRjsl42irUIJybQlk5GlR9ucM",
	"y": "7BKEPCB8gYjrx6p5DuMH0ckLyvn5stprWoUNdwW1L3k"
}`

const jwkFpKey2 = "0xc4db45d29987aead9b6fee307f5142a4af523b60"
const JwkKeyuidKey2 = "0x200000000000000000000000c4db45d29987aead9b6fee307f5142a4af523b60"

func TestECDSAFingerprint(t *testing.T) {
	fp, _ := ECDSAFingerprint(jwkPubKey1)
	assert.Equal(t, fp, jwkFpKey1, "no error")
	keyuid := ECDSAToB32(fp)
	ktype := B32Type(keyuid)
	assert.Equal(t, ktype, "ecdsa", "curve")
}

func TestECDSASign(t *testing.T) {
	// sig := ECDSASign(jwkPrivKey1, "message") //
	// VTyYt7K6ivCCiETREH6UwUSd4onQFLPZDN4zTrijvEA-jUYP_70NVxuwzYcX88ksgFt-HUaqxGzCSl5xo4Lp4w
	// HIzNxURC-mJDFMWGyoY_Qc1gp68gBdgG7LE7pOg15D3scy5iEhISw9tXYtbwH6FHo4jQ2rRhzdszhhkEQwv-nQ
	res, err := ECDSAVerify(jwkPubKey1, "message", "-32Xb5x06cc3vR3MLbCr51WhLjP3u6uJ3k9Gpq1_DHGD4i0grV1Sys8HKCspPJLZDdvMJiIjFCMaug-qRhaD2A")
	// res, err := ECDSAVerify(pub, "", "wDKeCt0NkKMAh69Z3U5Ix9LUBqbq-NzAGw2NlWjyrBuwhHYpS_bhD9_JfL9VCNHHnVclkYx_IAt_F46gOMNq1g")
	assert.Equal(t, err, nil, "no error expected")
	assert.Equal(t, res, true, "")
}

func TestECDSAVerify(t *testing.T) {
	sig, err := ECDSASign(jwkPrivKey1, "message")
	assert.Equal(t, err, nil, "no error expected")
	res, _ := ECDSAVerify(jwkPubKey1, "message", sig)
	assert.Equal(t, res, true, "")
}

func TestECDSAExportKey(t *testing.T) {
	buf := new(bytes.Buffer)
	json.Compact(buf, []byte(jwkPrivKey1))
	jwk := buf.String()
	k, _ := ECDSAExport(jwk)
	assert.Equal(t, k, jwkB64U, "")
}

func TestECDSAImport(t *testing.T) {
	k, _ := ECDSAExport(jwkPrivKey1)
	_, err := ECDSAImport(k)
	assert.Nil(t, err, "no error")
}
