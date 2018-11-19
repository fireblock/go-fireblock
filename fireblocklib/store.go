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

var stores map[string]*Store

// Store object
type Store struct {
	states map[string]interface{}
}

// NewStore initialize states
func NewStore(name string) *Store {
	if stores == nil {
		stores = make(map[string]*Store)
	}
	store := stores[name]
	if store == nil {
		st := new(Store)
		st.states = make(map[string]interface{})
		stores[name] = st
		return st
	}
	return store
}

// GetStore get store
func GetStore(name string) *Store {
	return stores[name]
}

// DelStore delete a store
func DelStore(name string) {
	delete(stores, name)
}

// GetString returns value or def if no value
func (store *Store) GetString(key, def string) string {
	val, ok := store.states[key]
	if ok {
		if v, ok := val.(string); ok {
			return v
		}
		return def
	}
	return def
}

// SetString sets or overwrites a value
func (store *Store) SetString(key, value string) {
	store.states[key] = value
}

// GetBool returns value or def if no value
func (store *Store) GetBool(key string, def bool) bool {
	val, ok := store.states[key]
	if ok {
		if v, ok := val.(bool); ok {
			return v
		}
		return def
	}
	return def
}

// SetBool sets or overwrites a value
func (store *Store) SetBool(key string, value bool) {
	store.states[key] = value
}
