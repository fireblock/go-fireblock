package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUrlContent1(t *testing.T) {
	res, err := getURLContent("https://twitter.com/fireblock_io/status/949233097060634624")
	assert.Equal(t, err, nil, "")
	assert.NotEqual(t, res, "", "not empty string")
}

func TestCheckTwitter1(t *testing.T) {
	r := CheckTwitter("https://twitter.com/ellis23232323/status/1016664114305921027", "ellis23232323", "rJpoNe2zm", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, r, true, "")
}

func TestGithub1(t *testing.T) {
	r := CheckGithub("https://gist.github.com/ellis2424/9397888371ed7b32d29d9e870cefd283", "ellis2424", "rJpoNe2zm", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, r, true, "")
}

func TestGithub2(t *testing.T) {
	r := CheckGithub("https://gist.github.com/ellis2424/ff72f8b5148dd372ae2f8b226f83f40c", "ellis2424", "S1H_x32Vf", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, r, false, "")
}

func TestCheckLinkedin(t *testing.T) {
	r := CheckLinkedin("https://www.linkedin.com/feed/update/urn:li:activity:6422410558120820736", "ellis-red-212223153", "rJpoNe2zm", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, r, true, "")
}

func TestCheckHTTPS(t *testing.T) {
	r := CheckHTTPS("https://dev.fireblock.io/.well-known/0x99090eae43316b2ba65ec52bcd5834a3e07edb2c.txt", "dev.fireblock.io", "rJpoNe2zm", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, r, true, "")
}
