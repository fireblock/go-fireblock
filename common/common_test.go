// common
package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	hash1 := Sha256("toto")
	assert.Equal(t, hash1, "0x31f7a65e315586ac198bd798b6629ce4903d0899476d5741a9f32e2e521b6a66", "they must be equal")
}

func TestIsSHA256(t *testing.T) {
	b1 := IsSha256("toto")
	assert.Equal(t, b1, false, "not a sha256 but a string")
	b2 := IsSha256("0x0")
	assert.Equal(t, b2, true, "a nul sha256")
	b3 := IsSha256("0x31f7a65e315586ac198bd798b6629ce4903d0899476d5741a9f32e2e521b6a66")
	assert.Equal(t, b3, true, "a real sha256")
	b4 := IsSha256("0x31f7a65e315586ac198bd798b6629ce4903d0899476d5741a9f32e2e521b6a6600")
	assert.Equal(t, b4, false, "too long for a sha256")
}
func TestSHA1(t *testing.T) {
	hash1 := Sha1("toto")
	assert.Equal(t, hash1, "0x0b9c2625dc21ef05f6ad4ddf47c5f203837aa32c", "they must be equal")
}
func TestKeccak256(t *testing.T) {
	hash1 := Keccak256("toto")
	assert.Equal(t, hash1, "0x2ef06b8bbad022ca2dd29795902ceb588d06d1cfd10cb6e687db0dbb837865e9", "they must be equal")
}

func TestPGPToB32(t *testing.T) {
	b32Key := PGPToB32("0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, len(b32Key), 66, "len is 66")
	assert.Equal(t, b32Key, "0x10000000000000000000000099090eae43316b2ba65ec52bcd5834a3e07edb2c")
}

func TestB32ToPGP(t *testing.T) {
	pgpKey := B32ToPGP("0x10000000000000000000000099090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, len(pgpKey), 42, "len is 42")
	assert.Equal(t, pgpKey, "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
}

func TestECDSAToB32(t *testing.T) {
	b32Key := ECDSAToB32("0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, len(b32Key), 66, "len is 66")
	assert.Equal(t, b32Key, "0x20000000000000000000000099090eae43316b2ba65ec52bcd5834a3e07edb2c")
}

func TestB32ToECDSA(t *testing.T) {
	pgpKey := B32ToECDSA("0x20000000000000000000000099090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, len(pgpKey), 42, "len is 42")
	assert.Equal(t, pgpKey, "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
}

func TestCheckB32Type(t *testing.T) {
	pgpType := B32Type("0x10000000000000000000000099090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, pgpType, "pgp", "this is a pgp key")
	ecdsaType := B32Type("0x20000000000000000000000099090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, ecdsaType, "ecdsa", "this is a ecdsa key")
	ethType := B32Type("0x00000000000000000000000099090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, ethType, "eth", "this is a eth key")
	unknownType := B32Type("0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, unknownType, "unknown", "this is an unknown eth key")
}
