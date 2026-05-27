// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package apptype

import (
	"sync"

	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

const (
	DEFAULT_MODULE        = "ace"
	DOC_MODULE            = "doc"
	TABLE_MODULE          = "table"
	PARAM_MODULE          = "param"
	APP                   = "app"
	HTML                  = "html"
	API                   = "api"
	PROXY                 = "proxy"
	FRAGMENT              = "fragment"
	STYLE                 = "style"
	REDIRECT              = "redirect"
	PERMISSION            = "permission"
	RESPONSE              = "response"
	CONFIG                = "config"
	LIBRARY               = "library"
	ACTION                = "action"
	RESULT                = "result"
	AUDIT                 = "audit"
	OUTPUT                = "output"
	CONTAINER_URL         = "<CONTAINER_URL>" // special url to use for proxying to the container
	DEFAULT_REDIRECT_CODE = 303
)

const (
	DEFAULT_DAISYUI_LIGHT_THEME = "emerald"
	DEFAULT_DAISYUI_DARK_THEME  = "night"
)

const (
	// Constants included in the ace builtin module
	GET       = "GET"
	POST      = "POST"
	PUT       = "PUT"
	DELETE    = "DELETE"
	HTML_TYPE = "HTML"
	JSON      = "JSON"
	TEXT      = "TEXT"
	READ      = "READ"
	WRITE     = "WRITE"

	AUTO     = "AUTO"
	TABLE    = "TABLE"
	DOWNLOAD = "DOWNLOAD"
	IMAGE    = "IMAGE"
)

var (
	once    sync.Once
	builtin starlark.StringDict
)

func createAppBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createHtmlBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createFragmentBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createStyleBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createRedirectBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createResponseBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createPermissionBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createLibraryBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createActionBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createResultBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createAuditBuiltin(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// Set the audit values in the thread local, that last call to audit from a handler takes effect

func createOutputBuiltin(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createProxyBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func createAPIBuiltin(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func CreateConfigBuiltin(nodeConfig types.NodeConfig, allowedEnv []string) func(_ *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return nil
}

// _branch is a special input that returns the current branch name

// _dev is a special input that returns the current dev status

// app_url is a special input that returns the current appUrl name

func CreateBuiltin(nodeConfig types.NodeConfig, allowedEnv []string) starlark.StringDict {
	_ = "STUB: not implemented"
	return *new(starlark.StringDict)
}
