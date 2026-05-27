// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package apptype

import (
	"go.starlark.net/starlark"
)

func GetStringAttr(s starlark.HasAttrs, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetOptionalStringAttr(s starlark.HasAttrs, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetIntAttr(s starlark.HasAttrs, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetBoolAttr(s starlark.HasAttrs, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetOptionalBoolAttr(s starlark.HasAttrs, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetListStringAttr(s starlark.HasAttrs, key string, optional bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetStringList(list *starlark.List) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetCallableAttr(s starlark.HasAttrs, key string) (starlark.Callable, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Callable), nil
}

func GetListMapAttr(s starlark.HasAttrs, key string, optional bool) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetDictAttr(s starlark.HasAttrs, key string, optional bool) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetListListStringAttr(s starlark.HasAttrs, key string, optional bool) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
