// Copyright 2015-2017 Fireblock.
// This file is part of Fireblock.

// Fireblock is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Fireblock is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Fireblock.  If not, see <http://www.gnu.org/licenses/>.

package fireblocklib

import (
	"testing"
)

var resultString string
var resultB32 [32]byte

func BenchmarkSha256(b *testing.B) {
	var res string
	for n := 0; n < b.N; n++ {
		res = Sha256("raw text")
	}
	resultString = res
}

// raw perf: 2.5Mi/s 1C 6400T
func BenchmarkSha256B(b *testing.B) {
	var res [32]byte
	var input = []byte("12345678123456781234567812345678")
	for n := 0; n < b.N; n++ {
		res = Sha256B(input)
	}
	resultB32 = res
}
