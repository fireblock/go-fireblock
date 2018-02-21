package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPKeyWithInvalidKeyuid(t *testing.T) {
	key, err := HTTPKey("0x0")
	e := err.(*FBKError)
	assert.Equal(t, e.Type(), InvalidKey, "no key found")
	assert.Equal(t, key, "", "no key found")
}

func TestHTTPKeyWithValidKeyuid(t *testing.T) {
	keyuid := PGPToB32("0x9ab6c2990618d54b3c0b3d56c55631b5d56e00fd")
	key, err := HTTPKey(keyuid)
	assert.Nil(t, err, "no error. key found")
	assert.Equal(t, len(key), 3141, "key found")
}

func TestHTTPCard(t *testing.T) {
	ServerURL = "https://dev.fireblock.io"
	keyuid := ECDSAToB32("0x3593fdb2f3ad14c024beaa389e4978a01ddd77d6")
	card, err := HTTPCard(keyuid, "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiIDogMTUxOTA1Njg4OCwgImV4cCI6IDE1NTA1OTI4ODgsICJ2ZXJpZmllZCI6IHRydWUsICJ1aWQiOiJsYXVyZW50Lm1hbGxldEBnbWFpbC5jb20iLCAidXVpZCI6ICIweDkxMjY3MzM2ZTdjNWI2M2I3NThhNDY1N2NmNTE1MTEwM2IxOWY2Zjg4NjE2MDlhMzk5ODY1MGI2NDFlYjQ1NmYiLCAia2V5dWlkIjogIjB4MjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMzU5M2ZkYjJmM2FkMTRjMDI0YmVhYTM4OWU0OTc4YTAxZGRkNzdkNiIgfQ.1ebrnK5w5uvb29YfJCQf6w6By8T")
	e := err.(*FBKError)
	assert.Equal(t, e.Type(), InvalidCard, "no card found")
	assert.Equal(t, card, "", "no key found")
}
