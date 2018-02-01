package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUrlContent1(t *testing.T) {
	res, err := getUrlContent("https://twitter.com/fireblock_io/status/949233097060634624")
	assert.Equal(t, err, nil, "")
	assert.NotEqual(t, res, "", "not empty string")
}

func TestCheckTwitter1(t *testing.T) {
	r := CheckTwitter("https://twitter.com/fireblock_io/status/949233097060634624", "fireblock_io", "HyPgw0hXf", "0xe204f3469808079909cafb1e4647f47e6a5c1742")
	assert.Equal(t, r, true, "")
}

func TestGithub1(t *testing.T) {
	r := CheckGithub("https://gist.github.com/ellis2424/435af0b890d1f465f29f16e19bcd8dc6", "ellis2424", "S1H_x32Vf", "0x3593fdb2f3ad14c024beaa389e4978a01ddd77d6")
	assert.Equal(t, r, true, "")
}

func TestGithub2(t *testing.T) {
	r := CheckGithub("https://gist.github.com/ellis2424/ff72f8b5148dd372ae2f8b226f83f40c", "ellis2424", "S1H_x32Vf", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, r, false, "")
}

func TestCheckLinkedin(t *testing.T) {
	r := CheckLinkedin("https://www.linkedin.com/feed/update/urn:li:activity:6337006284570787840", "ellis-red-212223153", "a4a23fbfd2", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, r, true, "")
}

func TestCheckHTTPS(t *testing.T) {
	r := CheckHTTPS("https://dev.fireblock.io/.fireblock/0x99090eae43316b2ba65ec52bcd5834a3e07edb2c", "dev.fireblock.io", "r1CI7XgzM", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
	assert.Equal(t, r, true, "")
}
