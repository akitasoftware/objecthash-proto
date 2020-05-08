// Copyright 2017 The ObjectHash-Proto Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protohash

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/fnv"
	"math"

	"github.com/OneOfOne/xxhash"
)

const (
	// Sorted alphabetically by value.
	boolIdentifier     = `b`
	mapIdentifier      = `d`
	floatIdentifier    = `f`
	intIdentifier      = `i`
	listIdentifier     = `l`
	nilIdentifier      = `n`
	byteIdentifier     = `r`
	unicodeIndentifier = `u`
)

type BasicHashFunc int

const (
	SHA256 BasicHashFunc = iota
	MD5
	FNV1A_128
	XXHASH64
)

func (f BasicHashFunc) String() string {
	switch f {
	case SHA256:
		return "SHA256"
	case MD5:
		return "MD5"
	case FNV1A_128:
		return "FNV-1a 128-bit"
	case XXHASH64:
		return "XXHASH64"
	default:
		return "UNKNOWN"
	}
}

type basicHasher struct {
	basicHashFunc BasicHashFunc
}

func (h basicHasher) hash(t string, b []byte) ([]byte, error) {
	var bh hash.Hash
	switch h.basicHashFunc {
	case SHA256:
		bh = sha256.New()
	case MD5:
		bh = md5.New()
	case FNV1A_128:
		bh = fnv.New128a()
	case XXHASH64:
		bh = xxhash.New64()
	default:
		return nil, fmt.Errorf("unsupported basic hash function: %d", h.basicHashFunc)
	}

	if _, err := bh.Write([]byte(t)); err != nil {
		return nil, err
	}

	if _, err := bh.Write(b); err != nil {
		return nil, err
	}

	return bh.Sum(nil), nil
}

func (h basicHasher) hashBool(b bool) ([]byte, error) {
	bb := []byte(`0`)
	if b {
		bb = []byte(`1`)
	}
	return h.hash(boolIdentifier, bb)
}

func (h basicHasher) hashBytes(bs []byte) ([]byte, error) {
	return h.hash(byteIdentifier, bs)
}

func (h basicHasher) hashFloat(f float64) ([]byte, error) {
	var normalizedFloat string

	switch {
	case math.IsInf(f, 1):
		normalizedFloat = "Infinity"
	case math.IsInf(f, -1):
		normalizedFloat = "-Infinity"
	case math.IsNaN(f):
		normalizedFloat = "NaN"
	default:
		var err error
		normalizedFloat, err = floatNormalize(f)
		if err != nil {
			return nil, err
		}
	}

	return h.hash(floatIdentifier, []byte(normalizedFloat))
}

func (h basicHasher) hashInt64(i int64) ([]byte, error) {
	return h.hash(intIdentifier, []byte(fmt.Sprintf("%d", i)))
}

func (h basicHasher) hashNil() ([]byte, error) {
	return h.hash(nilIdentifier, []byte(``))
}

func (h basicHasher) hashUint64(i uint64) ([]byte, error) {
	return h.hash(intIdentifier, []byte(fmt.Sprintf("%d", i)))
}

func (h basicHasher) hashUnicode(s string) ([]byte, error) {
	return h.hash(unicodeIndentifier, []byte(s))
}
