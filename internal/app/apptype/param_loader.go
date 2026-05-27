// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package apptype

import (
	"regexp"

	"github.com/openrundev/openrun/internal/app/starlark_type"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

const (
	PARAM = "param"
)

type DisplayType string

const (
	DisplayTypePassword   DisplayType = "password"
	DisplayTypeTextArea   DisplayType = "textarea"
	DisplayTypeFileUpload DisplayType = "file"
)

// AppParam represents a parameter in an app.
type AppParam struct {
	Index              int
	Name               string
	Description        string
	Required           bool
	Type               starlark_type.TypeName
	DefaultValue       starlark.Value
	DisplayType        DisplayType
	DisplayTypeOptions string
}

func ReadParamInfo(fileName string, inp []byte, serverConfig *types.ServerConfig) (map[string]AppParam, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var spaceRegex = regexp.MustCompile(`\s`)

func validateParamInfo(paramInfo map[string]AppParam) error { _ = "STUB: not implemented"; return nil }

func LoadParamInfo(fileName string, data []byte, serverConfig *types.ServerConfig) (map[string]AppParam, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: log

func ParamStringToType(name string, typeName starlark_type.TypeName, valueStr string) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}
