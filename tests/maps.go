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

package tests

import (
	"testing"

	"github.com/golang/protobuf/proto"

	oi "github.com/akitasoftware/objecthash-proto/internal"
	pb2_latest "github.com/akitasoftware/objecthash-proto/test_protos/generated/latest/proto2"
	pb3_latest "github.com/akitasoftware/objecthash-proto/test_protos/generated/latest/proto3"
	ti "github.com/akitasoftware/objecthash-proto/tests/internal"
)

// TestMaps performs tests on how maps are handled.
func TestMaps(t *testing.T, hashers oi.ProtoHashers) {
	hasher := hashers.StringPreferringHasher

	testCases := []ti.TestCase{
		////////////////////
		//  Boolean maps. //
		////////////////////
		{
			Protos: []proto.Message{
				&pb2_latest.BoolMaps{BoolToString: map[bool]string{true: "NOT FALSE", false: "NOT TRUE"}},
				&pb3_latest.BoolMaps{BoolToString: map[bool]string{true: "NOT FALSE", false: "NOT TRUE"}},
			},
			// No equivalent JSON object because JSON map keys must be strings.
			EquivalentObject:   map[string]map[bool]string{"bool_to_string": {true: "NOT FALSE", false: "NOT TRUE"}},
			ExpectedHashString: "d89d053bf7b37b4784832c72445661db99538fe1d490988575409a9040084f18",
		},

		////////////////////
		//  Integer maps. //
		////////////////////
		{
			Protos: []proto.Message{
				&pb2_latest.IntMaps{IntToString: map[int64]string{0: "ZERO"}},
				&pb3_latest.IntMaps{IntToString: map[int64]string{0: "ZERO"}},
			},
			// No equivalent JSON object because JSON map keys must be strings.
			EquivalentObject:   map[string]map[int64]string{"int_to_string": {0: "ZERO"}},
			ExpectedHashString: "53892192fb69cbd93ceb0552ca571b8505887f25d6f12822025341f16983a6af",
		},

		///////////////////
		//  String maps. //
		///////////////////
		{
			Protos: []proto.Message{
				&pb2_latest.StringMaps{StringToString: map[string]string{"foo": "bar"}},
				&pb3_latest.StringMaps{StringToString: map[string]string{"foo": "bar"}},
			},
			EquivalentJSONString: "{\"string_to_string\": {\"foo\": \"bar\"}}",
			EquivalentObject:     map[string]map[string]string{"string_to_string": {"foo": "bar"}},
			ExpectedHashString:   "cadfe560995647c63c20234a6409d2b1b8cf8dcf7d8e420ca33f23ff9ca9abfa",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.StringMaps{StringToString: map[string]string{
					"": "你好", "你好": "\u03d3", "\u03d3": "\u03d2\u0301"}},
				&pb3_latest.StringMaps{StringToString: map[string]string{
					"": "你好", "你好": "\u03d3", "\u03d3": "\u03d2\u0301"}},
			},
			EquivalentJSONString: "{\"string_to_string\": {\"\": \"你好\", \"你好\": \"\u03d3\", \"\u03d3\": \"\u03d2\u0301\"}}",
			EquivalentObject:     map[string]map[string]string{"string_to_string": {"": "你好", "你好": "\u03d3", "\u03d3": "\u03d2\u0301"}},
			ExpectedHashString:   "be8b5ae6d5986cde37ab8b395c66045fbb69a8b3b534fa34df7c19a640f4cd66",
		},

		//////////////////////////////
		//  Maps of proto messages. //
		//////////////////////////////
		{
			Protos: []proto.Message{
				&pb2_latest.StringMaps{StringToSimple: map[string]*pb2_latest.Simple{"foo": {}}},
				&pb3_latest.StringMaps{StringToSimple: map[string]*pb3_latest.Simple{"foo": {}}},
			},
			EquivalentJSONString: "{\"string_to_simple\": {\"foo\": {}}}",
			EquivalentObject:     map[string]map[string]map[string]string{"string_to_simple": {"foo": {}}},
			ExpectedHashString:   "58057927bb1a123452a2d75071b55b08e426490ca42c3dd14e3be59183ac4751",
		},
	}

	for _, tc := range testCases {
		tc.Check(t, hasher)
	}
}
