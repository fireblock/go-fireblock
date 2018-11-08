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

func TestCreateStore(t *testing.T) {
	store := NewStore("default")
	assert.NotNil(t, store, "create a store")
	DelStore("default")
}

func TestSetAndGet(t *testing.T) {
	store := NewStore("default")
	store.SetString("ellis", "my nickname")
	val := store.GetString("ellis", "")
	assert.Equal(t, val, "my nickname", "check set & get")
	val2 := store.GetString("ellis2", "default")
	assert.Equal(t, val2, "default", "check get default")
	val3 := store.GetString("ellis", "default")
	assert.Equal(t, val3, "my nickname", "check get value not default")
	DelStore("default")
}
