// Based on https://github.com/qri-io/starlib/blob/master/util/util.go
// Copyright (c) 2018 QRI, Inc. The MIT License (MIT)

package starlark_type

import (
	"go.starlark.net/starlark"
)

// UnquoteStarlark unquotes a starlark string value
func UnquoteStarlark(x starlark.Value) (string, error) { _ = "STUB: not implemented"; return "", nil }

// IsEmptyStarlarkString checks is a starlark string is empty ("" for a go string)
// starlark.String.String performs repr-style quotation, which is necessary
// for the starlark.Value contract but a frequent source of errors in API
// clients. This helper method makes sure it'll work properly
func IsEmptyStarlarkString(s starlark.String) bool { _ = "STUB: not implemented"; return false }

// UnmarshalStarlark decodes a starlark.Value into it's golang counterpart
func UnmarshalStarlark(x starlark.Value) (val interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// key as interface if found one key is not a string

// found key as not a string

// prepare result

// key as interface

// map[interface{}]interface{}

// map[string]interface{}

// MarshalStarlark turns go values into starlark types
func MarshalStarlark(data interface{}) (v starlark.Value, err error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// Unmarshaler is the interface use to unmarshal starlark custom types.
type Unmarshaler interface {
	// UnmarshalStarlark unmarshal a starlark object to custom type.
	UnmarshalStarlark(starlark.Value) error
}

// Marshaler is the interface use to marshal starlark custom types.
type Marshaler interface {
	// MarshalStarlark marshal a custom type to starlark object.
	MarshalStarlark() (starlark.Value, error)
}

// Unmarshaler is the interface use to unmarshal starlark custom types.
type TypeUnmarshaler interface {
	// UnmarshalStarlark unmarshals a starlark object to go object
	UnmarshalStarlarkType() (any, error)
}
