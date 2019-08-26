// Copyright 2018 The ObjectHash-Proto Authors
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
	ti "github.com/akitasoftware/objecthash-proto/tests/internal"
)

// TestProto2DefaultFieldValues checks that proto2 default field values are properly hashed.
func TestProto2DefaultFieldValues(t *testing.T, hashers oi.ProtoHashers) {
	hasher := hashers.StringPreferringHasher

	testCases := []ti.TestCase{
		{
			Protos: []proto.Message{
				&pb2_latest.Simple{BoolField: proto.Bool(false)},
			},
			EquivalentJSONString: "{\"bool_field\":false}",
			EquivalentObject:     map[string]bool{"bool_field": false},
			ExpectedHashString:   "1ab5ecdbe4176473024f7efd080593b740d22d076d06ea6edd8762992b484a12",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.Simple{BytesField: []byte{}},
			},
			// No JSON equivalent because JSON does not have a "bytes" representation.
			EquivalentObject:   map[string][]byte{"bytes_field": {}},
			ExpectedHashString: "10a0dbbfa097b731c7a505246ffa96a82f997b8c25892d76d3b8b1355e529e05",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.Simple{StringField: proto.String("")},
			},
			EquivalentJSONString: "{\"string_field\":\"\"}",
			EquivalentObject:     map[string]string{"string_field": ""},
			ExpectedHashString:   "2d60c2941830ef4bb14424e47c6cd010f2b95e5e34291f429998288a60ac8c22",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.Fixed32Message{Value: proto.Uint32(0)},
				&pb2_latest.Fixed64Message{Value: proto.Uint64(0)},
				&pb2_latest.Int32Message{Value: proto.Int32(0)},
				&pb2_latest.Int64Message{Value: proto.Int64(0)},
				&pb2_latest.Sfixed32Message{Value: proto.Int32(0)},
				&pb2_latest.Sfixed64Message{Value: proto.Int64(0)},
				&pb2_latest.Sint32Message{Value: proto.Int32(0)},
				&pb2_latest.Sint64Message{Value: proto.Int64(0)},
				&pb2_latest.Uint32Message{Value: proto.Uint32(0)},
				&pb2_latest.Uint64Message{Value: proto.Uint64(0)},
			},
			// No JSON equivalent because JSON does not have an integer representation.
			EquivalentObject:   map[string]int64{"value": 0},
			ExpectedHashString: "49f031b73dad26859ffeea8a2bb170aaf7358d2277b00c7fc7ea8edcd37e53a1",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(0.0)},
				&pb2_latest.FloatMessage{Value: proto.Float32(0.0)},
			},
			EquivalentJSONString: "{\"value\":0.0}",
			EquivalentObject:     map[string]float64{"value": 0.0},
			ExpectedHashString:   "94136b0850db069dfd7bee090fc7ede48aa7da53ae3cc8514140a493818c3b91",
		},
	}

	for _, tc := range testCases {
		tc.Check(t, hasher)
	}
}
