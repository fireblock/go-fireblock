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

func TestReadBatch(t *testing.T) {
	batch := `[{"h":"0x2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63993","f":"go1.11.1.linux-amd64.tar.gz","s":127205934,"t":"application/gzip"},{"h":"0xbac3c86a6a1591259e1e11fb4e94b015d17940b04280b9bb76a9626db573d9fd","f":"discord-0.0.5.deb","s":52100102,"t":"application/vnd.debian.binary-package"}]`
	res, err := readBatch(batch)
	assert.Equal(t, err, nil, "")
	assert.Equal(t, len(res), 2, "")
	assert.Equal(t, res[0].Hash, "0x2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63993", "")
}

func TestIsInBatch(t *testing.T) {
	batch := `[{"h":"0x2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63993","f":"go1.11.1.linux-amd64.tar.gz","s":127205934,"t":"application/gzip"},{"h":"0xbac3c86a6a1591259e1e11fb4e94b015d17940b04280b9bb76a9626db573d9fd","f":"discord-0.0.5.deb","s":52100102,"t":"application/vnd.debian.binary-package"}]`
	ck := IsInBatch(batch, "0x2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63993")
	assert.Equal(t, true, ck, "")
}
func TestNotInBatch(t *testing.T) {
	batch := `[{"h":"0x2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63992","f":"go1.11.1.linux-amd64.tar.gz","s":127205934,"t":"application/gzip"},{"h":"0xbac3c86a6a1591259e1e11fb4e94b015d17940b04280b9bb76a9626db573d9fd","f":"discord-0.0.5.deb","s":52100102,"t":"application/vnd.debian.binary-package"}]`
	ck := IsInBatch(batch, "0x2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63993")
	assert.Equal(t, false, ck, "")
}

func TestNotInBatchInvalidJSON(t *testing.T) {
	batch := `[{"h:"0x2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63992","f":"go1.11.1.linux-amd64.tar.gz","s":127205934,"t":"application/gzip"},{"h":"0xbac3c86a6a1591259e1e11fb4e94b015d17940b04280b9bb76a9626db573d9fd","f":"discord-0.0.5.deb","s":52100102,"t":"application/vnd.debian.binary-package"}]`
	ck := IsInBatch(batch, "0x2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63993")
	assert.Equal(t, false, ck, "")
}
