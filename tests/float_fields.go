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
	"math"
	"testing"

	"github.com/golang/protobuf/proto"

	oi "github.com/akitasoftware/objecthash-proto/internal"
	pb2_latest "github.com/akitasoftware/objecthash-proto/test_protos/generated/latest/proto2"
	pb3_latest "github.com/akitasoftware/objecthash-proto/test_protos/generated/latest/proto3"
	ti "github.com/akitasoftware/objecthash-proto/tests/internal"
)

// TestFloatFields performs tests on how floating point numbers are handled.
func TestFloatFields(t *testing.T, hashers oi.ProtoHashers) {
	hasher := hashers.StringPreferringHasher

	testCases := []ti.TestCase{
		/////////////////////////////
		//  Equivalence of Floats. //
		/////////////////////////////
		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Values: []float64{-2, -1, 0, 1, 2}},
				&pb3_latest.DoubleMessage{Values: []float64{-2, -1, 0, 1, 2}},

				&pb2_latest.FloatMessage{Values: []float32{-2, -1, 0, 1, 2}},
				&pb3_latest.FloatMessage{Values: []float32{-2, -1, 0, 1, 2}},
			},
			EquivalentObject:     map[string][]float64{"values": {-2, -1, 0, 1, 2}},
			EquivalentJSONString: "{\"values\": [-2, -1, 0, 1, 2]}",
			ExpectedHashString:   "586202dddb0e98bb8ce0b7289e29a9f7397b9b1996f3f8fe788f4cfb230b7ee8",
		},

		// Note that due to how floating point numbers work, we have to carefully
		// choose the values below in order for the decimal representation of the
		// test fractions to have 32-bit and 64-bit representations that are equal.
		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Values: []float64{0.0078125, 7.888609052210118e-31}},
				&pb3_latest.DoubleMessage{Values: []float64{0.0078125, 7.888609052210118e-31}},

				&pb2_latest.FloatMessage{Values: []float32{0.0078125, 7.888609052210118e-31}},
				&pb3_latest.FloatMessage{Values: []float32{0.0078125, 7.888609052210118e-31}},
			},
			EquivalentObject:     map[string][]float64{"values": {0.0078125, 7.888609052210118e-31}},
			EquivalentJSONString: "{\"values\": [0.0078125, 7.888609052210118e-31]}",
			ExpectedHashString:   "7b7cba0ed312bc6611f0523e7c46ce9a2ed9ecb798eb80e1cdf93c95faf503c7",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Values: []float64{-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},
				&pb3_latest.DoubleMessage{Values: []float64{-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},

				&pb2_latest.FloatMessage{Values: []float32{-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},
				&pb3_latest.FloatMessage{Values: []float32{-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},
			},
			EquivalentObject:     map[string][]float64{"values": {-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625}},
			EquivalentJSONString: "{\"values\": [-1.0, 1.5, 1000.000244140625, 1267650600228229401496703205376, 32.0, 13.0009765625]}",
			ExpectedHashString:   "ac261ff3d8b933998e3fea278539eb40b15811dd835d224e0150dce4794168b7",
		},

		/////////////////////////////////////////////////////////////////
		//  Non-equivalence of Floats using different representations. //
		/////////////////////////////////////////////////////////////////
		{
			Protos: []proto.Message{
				&pb2_latest.FloatMessage{Value: proto.Float32(0.1)},
				&pb3_latest.FloatMessage{Value: 0.1},

				// A float64 "0.1" is not equal to a float32 "0.1".
				// However, float32 "0.1" is equal to float64 "1.0000000149011612e-1".
				&pb2_latest.DoubleMessage{Value: proto.Float64(1.0000000149011612e-1)},
				&pb3_latest.DoubleMessage{Value: 1.0000000149011612e-1},
			},
			EquivalentObject:     map[string]float32{"value": 0.1},
			EquivalentJSONString: "{\"value\": 1.0000000149011612e-1}", // JSON objecthash only uses 64-bit floats.
			ExpectedHashString:   "7081ed6a1e7ad8e7f981a2894a3bd6d3b0b0033b69c03cce84b61dd063f4efaa",
		},

		// There's no float32 number that is equivalent to a float64 "0.1".
		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(0.1)},
				&pb3_latest.DoubleMessage{Value: 0.1},
			},
			EquivalentObject:     map[string]float64{"value": 0.1},
			EquivalentJSONString: "{\"value\": 0.1}",
			ExpectedHashString:   "e175fbe785bae88b598d3ecaad8a64d2a998e9f673173a226868f2ef312a5225",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.FloatMessage{Value: proto.Float32(1.2163543e+25)},
				&pb3_latest.FloatMessage{Value: 1.2163543e+25},

				// The decimal representation of the equivalent 64-bit float is different.
				&pb2_latest.DoubleMessage{Value: proto.Float64(1.2163543234531120e+25)},
				&pb3_latest.DoubleMessage{Value: 1.2163543234531120e+25},
			},
			EquivalentObject:     map[string]float32{"value": 1.2163543e+25},
			EquivalentJSONString: "{\"value\": 1.2163543234531120e+25}", // JSON objecthash only uses 64-bit floats.
			ExpectedHashString:   "bbb17cf7312f2ba5b0002d781f16d1ab50c3d25dc044ed3428750826a1c68653",
		},

		// There's no float32 number that is equivalent to a float64 "1e+25".
		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(1e+25)},
				&pb3_latest.DoubleMessage{Value: 1e+25},
			},
			EquivalentObject:     map[string]float64{"value": 1e+25},
			EquivalentJSONString: "{\"value\": 1e+25}",
			ExpectedHashString:   "874beabbede24974a9f3f74e3448670e0c42c0aaba082f18b963b72253649362",
		},

		//////////////////////
		//  Special values. //
		//////////////////////
		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(0)},
				&pb2_latest.FloatMessage{Value: proto.Float32(0)},
				// Proto3 zero values are indistinguishable from unset values.
			},
			EquivalentObject:     map[string]float64{"value": 0},
			EquivalentJSONString: "{\"value\":0}",
			ExpectedHashString:   "94136b0850db069dfd7bee090fc7ede48aa7da53ae3cc8514140a493818c3b91",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(math.NaN())},
				&pb3_latest.DoubleMessage{Value: math.NaN()},

				&pb2_latest.FloatMessage{Value: proto.Float32(float32(math.NaN()))},
				&pb3_latest.FloatMessage{Value: float32(math.NaN())},
			},
			EquivalentObject: map[string]float64{"value": math.NaN()},
			// No equivalent JSON: JSON does not support special float values.
			// See: https://tools.ietf.org/html/rfc4627#section-2.4
			ExpectedHashString: "16614de29b0823c41cabc993fa6c45da87e4e74c5d836edbcddcfaaf06ffafd1",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(math.Inf(1))},
				&pb3_latest.DoubleMessage{Value: math.Inf(1)},

				&pb2_latest.FloatMessage{Value: proto.Float32(float32(math.Inf(1)))},
				&pb3_latest.FloatMessage{Value: float32(math.Inf(1))},
			},
			EquivalentObject: map[string]float64{"value": math.Inf(1)},
			// No equivalent JSON: JSON does not support special float values.
			// See: https://tools.ietf.org/html/rfc4627#section-2.4
			ExpectedHashString: "c58cd512e86204e99cb6c11d83bb3daaccdd946e66383004cb9b7f87f762935c",
		},

		{
			Protos: []proto.Message{
				&pb2_latest.DoubleMessage{Value: proto.Float64(math.Inf(-1))},
				&pb3_latest.DoubleMessage{Value: math.Inf(-1)},

				&pb2_latest.FloatMessage{Value: proto.Float32(float32(math.Inf(-1)))},
				&pb3_latest.FloatMessage{Value: float32(math.Inf(-1))},
			},
			EquivalentObject: map[string]float64{"value": math.Inf(-1)},
			// No equivalent JSON: JSON does not support special float values.
			// See: https://tools.ietf.org/html/rfc4627#section-2.4
			ExpectedHashString: "1a4ffd7e9dc1f915c5b3b821d9194ac7d6d2bdec947aa8c3b3b1e9017c651331",
		},
	}

	for _, tc := range testCases {
		tc.Check(t, hasher)
	}
}
