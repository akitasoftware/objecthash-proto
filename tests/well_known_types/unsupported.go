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
	"testing"

	"github.com/golang/protobuf/proto"
	any_pb "github.com/golang/protobuf/ptypes/any"
	duration_pb "github.com/golang/protobuf/ptypes/duration"
	struct_pb "github.com/golang/protobuf/ptypes/struct"
	wrappers_pb "github.com/golang/protobuf/ptypes/wrappers"

	oi "github.com/akitasoftware/objecthash-proto/internal"
	custom "github.com/akitasoftware/objecthash-proto/test_protos/custom"
	pb2_latest "github.com/akitasoftware/objecthash-proto/test_protos/generated/latest/proto2"
	pb3_latest "github.com/akitasoftware/objecthash-proto/test_protos/generated/latest/proto3"
)

// TestUnsupportedWellKnownTypes confirms that hashing any of the unsupported
// well-known types returns an error.
//
// Once support is added for any of those fields, a separate test method will
// be added for it.
func TestUnsupportedWellKnownTypes(t *testing.T, hashers oi.ProtoHashers) {
	hasher := hashers.DefaultHasher

	unsupportedProtos := []proto.Message{
		&any_pb.Any{},
		&pb2_latest.KnownTypes{AnyField: &any_pb.Any{}},
		&pb3_latest.KnownTypes{AnyField: &any_pb.Any{}},

		&wrappers_pb.BoolValue{},
		&pb2_latest.KnownTypes{BoolValueField: &wrappers_pb.BoolValue{}},
		&pb3_latest.KnownTypes{BoolValueField: &wrappers_pb.BoolValue{}},

		&wrappers_pb.BytesValue{},
		&pb2_latest.KnownTypes{BytesValueField: &wrappers_pb.BytesValue{}},
		&pb3_latest.KnownTypes{BytesValueField: &wrappers_pb.BytesValue{}},

		&wrappers_pb.DoubleValue{},
		&pb2_latest.KnownTypes{DoubleValueField: &wrappers_pb.DoubleValue{}},
		&pb3_latest.KnownTypes{DoubleValueField: &wrappers_pb.DoubleValue{}},

		&duration_pb.Duration{},
		&pb2_latest.KnownTypes{DurationField: &duration_pb.Duration{}},
		&pb3_latest.KnownTypes{DurationField: &duration_pb.Duration{}},

		&wrappers_pb.FloatValue{},
		&pb2_latest.KnownTypes{FloatValueField: &wrappers_pb.FloatValue{}},
		&pb3_latest.KnownTypes{FloatValueField: &wrappers_pb.FloatValue{}},

		&wrappers_pb.Int32Value{},
		&pb2_latest.KnownTypes{Int32ValueField: &wrappers_pb.Int32Value{}},
		&pb3_latest.KnownTypes{Int32ValueField: &wrappers_pb.Int32Value{}},

		&wrappers_pb.Int64Value{},
		&pb2_latest.KnownTypes{Int64ValueField: &wrappers_pb.Int64Value{}},
		&pb3_latest.KnownTypes{Int64ValueField: &wrappers_pb.Int64Value{}},

		&struct_pb.ListValue{},
		&pb2_latest.KnownTypes{ListValueField: &struct_pb.ListValue{}},
		&pb3_latest.KnownTypes{ListValueField: &struct_pb.ListValue{}},

		&wrappers_pb.StringValue{},
		&pb2_latest.KnownTypes{StringValueField: &wrappers_pb.StringValue{}},
		&pb3_latest.KnownTypes{StringValueField: &wrappers_pb.StringValue{}},

		&struct_pb.Struct{},
		&pb2_latest.KnownTypes{StructField: &struct_pb.Struct{}},
		&pb3_latest.KnownTypes{StructField: &struct_pb.Struct{}},

		&wrappers_pb.UInt32Value{},
		&pb2_latest.KnownTypes{Uint32ValueField: &wrappers_pb.UInt32Value{}},
		&pb3_latest.KnownTypes{Uint32ValueField: &wrappers_pb.UInt32Value{}},

		&wrappers_pb.UInt64Value{},
		&pb2_latest.KnownTypes{Uint64ValueField: &wrappers_pb.UInt64Value{}},
		&pb3_latest.KnownTypes{Uint64ValueField: &wrappers_pb.UInt64Value{}},

		&struct_pb.Value{},
		&pb2_latest.KnownTypes{ValueField: &struct_pb.Value{}},
		&pb3_latest.KnownTypes{ValueField: &struct_pb.Value{}},

		// Check that a future well-known type is unsupported by default.
		&custom.FutureWellKnownType{},
	}
	for _, message := range unsupportedProtos {
		_, err := hasher.HashProto(message)
		if err == nil {
			t.Errorf("Attempting to hash %T{ %+v} should have returned an error.", message, message)
		}
	}
}
