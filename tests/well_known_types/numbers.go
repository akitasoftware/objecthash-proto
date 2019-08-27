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

package wellknowntypes

import (
	"math"
	"testing"

	"github.com/golang/protobuf/proto"
	wrappers_pb "github.com/golang/protobuf/ptypes/wrappers"

	oi "github.com/akitasoftware/objecthash-proto/internal"
	pb2_latest "github.com/akitasoftware/objecthash-proto/test_protos/generated/latest/proto2"
	pb3_latest "github.com/akitasoftware/objecthash-proto/test_protos/generated/latest/proto3"
	ti "github.com/akitasoftware/objecthash-proto/tests/internal"
)

// TestNumericWellKnownTypes confirms that google.protobuf.Int32Value,
// google.protobuf.Int64Value, google.protobuf.UInt32Value,
// google.protobuf.UInt64Value, google.protobuf.FloatValue, and
// google.protobuf.DoubleValue protos are hashed properly.
func TestNumericWellKnownTypes(t *testing.T, hashers oi.ProtoHashers) {
	hasher := hashers.FieldNamesAsKeysHasher

	testCases := []ti.TestCase{
		//////////////////////////////
		//  Empty/Zero Wrappers. //
		//////////////////////////////

		// The semantics of these wrapper objects imply that the distinction
		// between unset and zero happen at the message level, rather than the
		// field level.
		//
		// As a result, an unset int32/int64/uint32/uint64/float/double is one
		// where the proto itself is nil, while an explicitly set proto with
		// unset fields is considered to be explicitly set to 0.
		//
		// This is unlike normal proto3 messages, where unset/zero fields must be
		// considered to be unset, because they're indistinguishable in the general
		// case.
		{
			Protos: []proto.Message{
				&wrappers_pb.Int32Value{},
				&wrappers_pb.Int32Value{Value: 0},
			},
			EquivalentObject:   int32(0),
			ExpectedHashString: "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.Int64Value{},
				&wrappers_pb.Int64Value{Value: 0},
			},
			EquivalentObject:   int64(0),
			ExpectedHashString: "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.UInt32Value{},
				&wrappers_pb.UInt32Value{Value: 0},
			},
			EquivalentObject:   uint32(0),
			ExpectedHashString: "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.UInt64Value{},
				&wrappers_pb.UInt64Value{Value: 0},
			},
			EquivalentObject:   uint64(0),
			ExpectedHashString: "a4e167a76a05add8a8654c169b07b0447a916035aef602df103e8ae0fe2ff390",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.FloatValue{},
				&wrappers_pb.FloatValue{Value: 0},
			},
			EquivalentObject:   float32(0),
			ExpectedHashString: "60101d8c9cb988411468e38909571f357daa67bff5a7b0a3f9ae295cd4aba33d",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.DoubleValue{},
				&wrappers_pb.DoubleValue{Value: 0},
			},
			EquivalentObject:   float64(0),
			ExpectedHashString: "60101d8c9cb988411468e38909571f357daa67bff5a7b0a3f9ae295cd4aba33d",
		},

		/////////////////////////
		//  Normal Wrappers. //
		/////////////////////////
		{
			Protos: []proto.Message{
				&wrappers_pb.Int32Value{Value: math.MaxInt32},
			},
			EquivalentObject:   int32(math.MaxInt32),
			ExpectedHashString: "4c46d595c28a829ed91f8feee378e34665f5b3f5cd0f35eb2e3ef115e96eef4f",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.Int64Value{Value: math.MaxInt64},
			},
			EquivalentObject:   int64(math.MaxInt64),
			ExpectedHashString: "eda7a99bc44462f5181f63a88e2ab9d8d318d68c2c2bf9ff70d9e4ecc2a99df3",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.UInt32Value{Value: math.MaxUint32},
			},
			EquivalentObject:   uint32(math.MaxUint32),
			ExpectedHashString: "88cdf1c5990befa03b32701a290ecbf7da4df8affaa3a12fcda66b23da3643fd",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.UInt64Value{Value: math.MaxUint64},
			},
			EquivalentObject:   uint64(math.MaxUint64),
			ExpectedHashString: "5b50a7751238c21772625d9807fc62e2d25ae5bd092d2018f0834d871c5db302",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.FloatValue{Value: math.MaxFloat32},
			},
			EquivalentObject:   float32(math.MaxFloat32),
			ExpectedHashString: "31ca3114782b94b13f9b299a9ea60c1db0c81ebf3474954ce7a8c5c22d408a1d",
		},
		{
			Protos: []proto.Message{
				&wrappers_pb.DoubleValue{Value: math.MaxFloat64},
			},
			EquivalentObject:   math.MaxFloat64,
			ExpectedHashString: "cb3a4a934c9e971271c4a5ce3987fdf7cecdbe7087c19496c4f7dceea6e74301",
		},

		//////////////////////////////////////
		//  Wrappers within other protos. //
		//////////////////////////////////////

		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{Int32ValueField: &wrappers_pb.Int32Value{}},
				&pb2_latest.KnownTypes{Int32ValueField: &wrappers_pb.Int32Value{Value: 0}},
				&pb3_latest.KnownTypes{Int32ValueField: &wrappers_pb.Int32Value{}},
				&pb3_latest.KnownTypes{Int32ValueField: &wrappers_pb.Int32Value{Value: 0}},
			},
			EquivalentObject:   map[string]int32{"int32_value_field": 0},
			ExpectedHashString: "f45c9b89d9a758f70fee58bad947bca07bd20a31119d927588e7bb11ef17180d",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{Int32ValueField: &wrappers_pb.Int32Value{Value: math.MaxInt32}},
				&pb3_latest.KnownTypes{Int32ValueField: &wrappers_pb.Int32Value{Value: math.MaxInt32}},
			},
			EquivalentObject:   map[string]int32{"int32_value_field": math.MaxInt32},
			ExpectedHashString: "b1621e15db55e9bccb00d48d24590b92b53758c1488336dec64c7a6422e9edcd",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{Int64ValueField: &wrappers_pb.Int64Value{}},
				&pb2_latest.KnownTypes{Int64ValueField: &wrappers_pb.Int64Value{Value: 0}},
				&pb3_latest.KnownTypes{Int64ValueField: &wrappers_pb.Int64Value{}},
				&pb3_latest.KnownTypes{Int64ValueField: &wrappers_pb.Int64Value{Value: 0}},
			},
			EquivalentObject:   map[string]int64{"int64_value_field": 0},
			ExpectedHashString: "8459ba1e83e7c72aeb9dcb564daf945f42fe3c1b8837b4266fac7754657160a1",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{Int64ValueField: &wrappers_pb.Int64Value{Value: math.MaxInt64}},
				&pb3_latest.KnownTypes{Int64ValueField: &wrappers_pb.Int64Value{Value: math.MaxInt64}},
			},
			EquivalentObject:   map[string]int64{"int64_value_field": math.MaxInt64},
			ExpectedHashString: "50110e3d2474a0c611da8d3f0565459fedaef3ef5b6829d9a6042c58854ec2a7",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint32ValueField: &wrappers_pb.UInt32Value{}},
				&pb2_latest.KnownTypes{Uint32ValueField: &wrappers_pb.UInt32Value{Value: 0}},
				&pb3_latest.KnownTypes{Uint32ValueField: &wrappers_pb.UInt32Value{}},
				&pb3_latest.KnownTypes{Uint32ValueField: &wrappers_pb.UInt32Value{Value: 0}},
			},
			EquivalentObject:   map[string]uint32{"uint32_value_field": 0},
			ExpectedHashString: "7e3d86d713dec0db2344ff4eb01e40b4cc2c8393840422cf6a716f220b6f6b69",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint32ValueField: &wrappers_pb.UInt32Value{Value: math.MaxUint32}},
				&pb3_latest.KnownTypes{Uint32ValueField: &wrappers_pb.UInt32Value{Value: math.MaxUint32}},
			},
			EquivalentObject:   map[string]uint32{"uint32_value_field": math.MaxUint32},
			ExpectedHashString: "aa86043990f6dddd1d8bb5e144d357d494e7071065a7984159f9c2f53f3c1225",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint64ValueField: &wrappers_pb.UInt64Value{}},
				&pb2_latest.KnownTypes{Uint64ValueField: &wrappers_pb.UInt64Value{Value: 0}},
				&pb3_latest.KnownTypes{Uint64ValueField: &wrappers_pb.UInt64Value{}},
				&pb3_latest.KnownTypes{Uint64ValueField: &wrappers_pb.UInt64Value{Value: 0}},
			},
			EquivalentObject:   map[string]uint64{"uint64_value_field": 0},
			ExpectedHashString: "832f86706cc1b4136e174c5f0814e965388b01ecad751f1bd23c7523a684b1cc",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{Uint64ValueField: &wrappers_pb.UInt64Value{Value: math.MaxUint64}},
				&pb3_latest.KnownTypes{Uint64ValueField: &wrappers_pb.UInt64Value{Value: math.MaxUint64}},
			},
			EquivalentObject:   map[string]uint64{"uint64_value_field": math.MaxUint64},
			ExpectedHashString: "ac227c7300873771ea3582d01b70e1e33a32bc801a28aad304db806a11c4432a",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{FloatValueField: &wrappers_pb.FloatValue{}},
				&pb2_latest.KnownTypes{FloatValueField: &wrappers_pb.FloatValue{Value: 0}},
				&pb3_latest.KnownTypes{FloatValueField: &wrappers_pb.FloatValue{}},
				&pb3_latest.KnownTypes{FloatValueField: &wrappers_pb.FloatValue{Value: 0}},
			},
			EquivalentObject:   map[string]float32{"float_value_field": 0},
			ExpectedHashString: "75085520c0294c8467895b2bd9939cf4a6373629f95f155eb3c755c7debb326d",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{FloatValueField: &wrappers_pb.FloatValue{Value: math.MaxFloat32}},
				&pb3_latest.KnownTypes{FloatValueField: &wrappers_pb.FloatValue{Value: math.MaxFloat32}},
			},
			EquivalentObject:   map[string]float32{"float_value_field": math.MaxFloat32},
			ExpectedHashString: "73808d9759e7494e379ec6f739f2728f51befd6caa86efcaa8ff550fc173d2fc",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{DoubleValueField: &wrappers_pb.DoubleValue{}},
				&pb2_latest.KnownTypes{DoubleValueField: &wrappers_pb.DoubleValue{Value: 0}},
				&pb3_latest.KnownTypes{DoubleValueField: &wrappers_pb.DoubleValue{}},
				&pb3_latest.KnownTypes{DoubleValueField: &wrappers_pb.DoubleValue{Value: 0}},
			},
			EquivalentObject:   map[string]float64{"double_value_field": 0},
			ExpectedHashString: "d593d09e840e41b2f5169561acf24a6b094f0dfb6850cf2a6dcea612f8990a41",
		},
		{
			Protos: []proto.Message{
				&pb2_latest.KnownTypes{DoubleValueField: &wrappers_pb.DoubleValue{Value: math.MaxFloat64}},
				&pb3_latest.KnownTypes{DoubleValueField: &wrappers_pb.DoubleValue{Value: math.MaxFloat64}},
			},
			EquivalentObject:   map[string]float64{"double_value_field": math.MaxFloat64},
			ExpectedHashString: "442120b4256374165fe184eac3db1fdf3b200ebb32777c0e936893e8e0c3de2a",
		},
	}

	for _, tc := range testCases {
		tc.Check(t, hasher)
	}
}
