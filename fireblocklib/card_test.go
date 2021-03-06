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

func TestCheckCard1(t *testing.T) {
	card := `[
{ "uid": "rJpoNe2zm", "provider": "fireblock", "date": "20180105" },
{ "uid": "ellis2424", "provider": "github", "proof": "https://gist.github.com/ellis2424/b241a2cfce6b5bfbfd3ab06b551171b3", "date": "20180116" },
{ "uid": "ellis23232323", "provider": "twitter", "proof": "https://twitter.com/ellis23232323/status/1016664114305921027", "date": "20180116" }, 
{ "uid": "dev.fireblock.io", "provider": "https", "proof": "https://dev.fireblock.io/.well-known/0x99090eae43316b2ba65ec52bcd5834a3e07edb2c.txt", "date": "20180116" },
{ "uid": "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c", "provider": "pgp", "date": "20180116" }
]`
	_, _, _, err := CheckCard(card, "0xf2e878d56dc604154ab07e5700e28da611b606b73ad5d75041ed5e5afbf699be")
	assert.Equal(t, err, nil, "")
}

func TestCheckCardError(t *testing.T) {
	_, _, _, err := CheckCard("{ invalid json", "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.NotNil(t, err, "invalid json")
	_, _, _, err = CheckCard(`[ { "uid": "dd", "provider": "twitter", "date": "20180105" } ]`, "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.NotNil(t, err, "missing fireblock element")
	_, _, _, err = CheckCard(`[ { "uid": "dd", "provider": "fireblock", "date": "20180105" } ]`, "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.NotNil(t, err, "bad fireblock element")
	_, _, _, err = CheckCard("[ { \"uid\": \"HyPgw0hXf\", \"provider\": \"fireblock\", \"date\": \"20180105\" }, { \"uid\": \"fireblockdev\", \"provider\": \"github\", \"proof\": \"https://gist.github.com/fireblockdev/de9fdfca46e6d707b99ddd7043b21ec6\", \"date\": \"20180116\" }, { \"uid\": \"fireblock_io\", \"provider\": \"twitter\", \"proof\": \"https://twitter.com/fireblock_io/status/949233097060634624\", \"date\": \"20180116\" }, { \"uid\": \"fireblock.io\", \"provider\": \"https\", \"proof\": \"https://fireblock.io/.fireblock/0xe204f3469808079909cafb1e4647f47e6a5c1742\", \"date\": \"20180116\" }, { \"uid\": \"0xe204f3469808079909cafb1e4647f47e6a5c1742\", \"provider\": \"pgp\" } ]", "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.NotNil(t, err, "error")
}

/* func TestCheckAllCard1(t *testing.T) {
	card := `[
{ "uid": "rJpoNe2zm", "provider": "fireblock", "date": "20180105" },
{ "uid": "ellis2424", "provider": "github", "proof": "https://gist.github.com/ellis2424/9397888371ed7b32d29d9e870cefd283", "date": "20180116" },
{ "uid": "ellis23232323", "provider": "twitter", "proof": "https://twitter.com/ellis23232323/status/1016664114305921027", "date": "20180116" },
{ "uid": "dev.fireblock.io", "provider": "https", "proof": "https://dev.fireblock.io/.well-known/0x99090eae43316b2ba65ec52bcd5834a3e07edb2c.txt", "date": "20180116" },
{ "uid": "0x99090eae43316b2ba65ec52bcd5834a3e07edb2c", "provider": "pgp", "date": "20180116" }
]`
	_, _, _, err := CheckAllCard(card, "0x4fd0a4896f9521a87e58d8c4ba45340b6d0aa19f5f6f3599a20e4f608f324592")
	assert.Equal(t, err, nil, "")
} */

func TestCheckAllCardError(t *testing.T) {
	_, _, _, err := CheckAllCard("{ invalid json", "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.NotNil(t, err, "invalid json")
	_, _, _, err = CheckAllCard(`[ { "uid": "dd", "provider": "twitter", "date": "20180105" } ]`, "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.NotNil(t, err, "missing fireblock element")
	_, _, _, err = CheckAllCard(`[ { "uid": "dd", "provider": "fireblock", "date": "20180105" } ]`, "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.NotNil(t, err, "bad fireblock element")
	_, _, _, err = CheckAllCard("[ { \"uid\": \"HyPgw0hXf\", \"provider\": \"fireblock\", \"date\": \"20180105\" }, { \"uid\": \"fireblockdev\", \"provider\": \"github\", \"proof\": \"https://gist.github.com/fireblockdev/de9fdfca46e6d707b99ddd7043b21ec6\", \"date\": \"20180116\" }, { \"uid\": \"fireblock_io\", \"provider\": \"twitter\", \"proof\": \"https://twitter.com/fireblock_io/status/949233097060634624\", \"date\": \"20180116\" }, { \"uid\": \"fireblock.io\", \"provider\": \"https\", \"proof\": \"https://fireblock.io/.fireblock/0xe204f3469808079909cafb1e4647f47e6a5c1742\", \"date\": \"20180116\" }, { \"uid\": \"0xe204f3469808079909cafb1e4647f47e6a5c1742\", \"provider\": \"pgp\" } ]", "0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7")
	assert.NotNil(t, err, "error")
}
