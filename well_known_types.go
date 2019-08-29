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

package protohash

import (
	"bytes"
	"fmt"
	"reflect"
)

// Supported well-known types.
var (
	timestamp string   = "Timestamp"
	numbers   []string = []string{
		"Int32Value",
		"Int64Value",
		"UInt32Value",
		"UInt64Value",
		"DoubleValue",
		"FloatValue",
	}
)

// hashWellKnownType hashes proto messages that are Well-known types.
//
// This method uses the reflect.Value of a well-known type's underlying struct
// object to calculate its hash.
//
// Well-known types are proto messages that have special semantics which are
// defined within the proto library. As a result, special treatment while
// calculating their hash is often (but not always) needed.
func (hasher *objectHasher) hashWellKnownType(name string, sv reflect.Value) ([]byte, error) {
	if name == timestamp {
		return hasher.hashTimestamp(sv)
	}
	for _, numName := range numbers {
		if name == numName {
			return hasher.hashNumber(sv)
		}
	}

	return nil, fmt.Errorf("Got a currently unsupported protobuf well-known type: %s", name)
}

// hashTimestamp calculates the object hash of a google.protobuf.Timestamp.
//
// This will be equivalent to the ObjectHash of a list of two integers, where
// the first list item is a non-negative integer equal to the timestamp's UTC
// seconds since epoch, and the second list item is a non-negative integer
// equal to the timestamp's fractions of a second at nanosecond resolution.
//
// Additionally, the semantics of the Timestamp object imply that the
// distinction between unset and zero happen at the message level, rather than
// the field level.  As a result, an unset timestamp is one where the proto
// itself is nil, while an explicitly set timestamp with unset fields is
// considered to be explicitly set to 0.  This is unlike normal proto3
// messages, where unset/zero fields must be considered to be unset, because
// they're indistinguishable in the general case.
//
// Note that this function's argument is a reflect.Value of the underlying
// struct object, rather than the proto message itself.
func (hasher *objectHasher) hashTimestamp(sv reflect.Value) ([]byte, error) {
	sk := sv.Kind()
	if sk != reflect.Struct {
		return nil, fmt.Errorf("Got a bad google.protobuf.Timestamp proto: %v. Expected a Struct, instead got a %s", sv, sk)
	}

	b := new(bytes.Buffer)

	// Hash seconds and nanoseconds.
	for _, field := range []string{"Seconds", "Nanos"} {
		fieldValue := sv.FieldByName(field)
		fk := fieldValue.Kind()
		if fk != reflect.Int64 && fk != reflect.Int32 {
			return nil, fmt.Errorf("Got a google.protobuf.Timestamp proto with a bad '%s' field: %v. Expected an integer, instead got a %s", field, sv, fk)
		}
		h, err := hasher.basicHasher.hashInt64(fieldValue.Int())
		if err != nil {
			return nil, err
		}
		b.Write(h[:])
	}

	return hasher.basicHasher.hash(listIdentifier, b.Bytes())
}

// hashNumber calculates the object hash of a google.protobuf.Int32Value,
// google.protobuf.Int64Value, google.protobuf.UInt32Value,
// google.protobuf.UInt64Value, google.protobuf.FloatValue, or
// google.protobuf.DoubleValue.
//
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
//
// Note that this function's argument is a reflect.Value of the underlying
// struct object, rather than the proto message itself.
func (hasher *objectHasher) hashNumber(sv reflect.Value) ([]byte, error) {
	sk := sv.Kind()
	if sk != reflect.Struct {
		return nil, fmt.Errorf("Got a bad google.protobuf.Int32Value/Int64Value/UInt32Value/UInt64Value/DoubleValue/FloatValue proto: %v. Expected a Struct, instead got a %s", sv, sk)
	}

	// Hash seconds and nanoseconds.
	v := sv.FieldByName("Value")
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return hasher.basicHasher.hashFloat(v.Float())
	case reflect.Int32, reflect.Int64:
		return hasher.basicHasher.hashInt64(v.Int())
	case reflect.Uint32, reflect.Uint64:
		return hasher.basicHasher.hashUint64(v.Uint())
	default:
		return nil, fmt.Errorf("Unsupported value type in well-known numeric protobuf: %T", v)
	}
}
