// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package apptype

import (
	"github.com/openrundev/openrun/internal/app/starlark_type"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

const (
	TYPE  = "type"
	FIELD = "field"
	INDEX = "index"
)

func ReadStoreInfo(fileName string, inp []byte) (*starlark_type.StoreInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateStoreInfo(storeInfo *starlark_type.StoreInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func LoadStoreInfo(fileName string, data []byte) (*starlark_type.StoreInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: log

func createStoreInfo(definedTypes map[string]starlark.Value, data []byte) (*starlark_type.StoreInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFields(typeName string, typeStruct *starlarkstruct.Struct, key string) ([]starlark_type.StoreField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attr is present

func getIndexes(typeName string, typeStruct *starlarkstruct.Struct, key string) ([]starlark_type.Index, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no indexes

func createFieldBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createIndexBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}
