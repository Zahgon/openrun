// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package starlark_type

import (
	"go.starlark.net/starlark"
)

type TypeName string

const (
	INT      TypeName = "INT"
	FLOAT    TypeName = "FLOAT"
	DATETIME TypeName = "DATETIME"
	STRING   TypeName = "STRING"
	BOOLEAN  TypeName = "BOOLEAN"
	LIST     TypeName = "LIST"
	DICT     TypeName = "DICT"
)

type StoreInfo struct {
	Bytes []byte // The raw bytes for the schema.star file
	Types []StoreType
}

type StoreType struct {
	Name    string
	Fields  []StoreField
	Indexes []Index
}

type StoreField struct {
	Name    string
	Type    TypeName
	Default any
}

type Index struct {
	Fields []string
	Unique bool
}

type TypeBuilder struct {
	Name   string
	Fields []StoreField
}

func (s *TypeBuilder) CreateType(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// Unpack takes field name followed by a pointer to the value

// Add value pointer
