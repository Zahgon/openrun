// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package starlark_type

import (
	"go.starlark.net/starlark"
)

// StarlarkType represents a Starlark type created from the schema type definition.
type StarlarkType struct {
	name string
	data map[string]starlark.Value
	keys []string
}

var _ starlark.Value = (*StarlarkType)(nil)

func NewStarlarkType(name string, data map[string]starlark.Value) *StarlarkType {
	_ = "STUB: not implemented"
	return nil
}

func (s *StarlarkType) Attr(attr string) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *StarlarkType) AttrNames() []string { _ = "STUB: not implemented"; return nil }

func (s *StarlarkType) SetField(name string, val starlark.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *StarlarkType) String() string { _ = "STUB: not implemented"; return "" }

func (s *StarlarkType) Type() string { _ = "STUB: not implemented"; return "" }

func (s *StarlarkType) Freeze() {
	_ = "STUB: not implemented"
	// Not supported
	return
}

func (s *StarlarkType) Truth() starlark.Bool { _ = "STUB: not implemented"; return *new(starlark.Bool) }

func (s *StarlarkType) Hash() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *StarlarkType) UnmarshalStarlarkType() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// ConvertToMap converts a struct to a map[string]any
func ConvertToMap(p any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

func ConvertToStarlark(p any) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}
