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
	r := CheckGithub("https://gist.github.com/ellis2424/b241a2cfce6b5bfbfd3ab06b551171b3", "ellis2424", "rJpoNe2zm", "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c")
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
