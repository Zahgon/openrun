// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"github.com/openrundev/openrun/internal/app/starlark_type"
	"go.starlark.net/starlark"
)

// PluginResponse is a starlark.Value that represents the response to a plugin request
type PluginResponse struct {
	errorCode int
	err       error
	value     any
	isStream  bool
	thread    *starlark.Thread
}

func NewErrorResponse(err error, thread *starlark.Thread) *PluginResponse {
	_ = "STUB: not implemented"
	return nil
}

func NewErrorCodeResponse(errorCode int, err error, value any) *PluginResponse {
	_ = "STUB: not implemented"
	return nil
}

func NewResponse(value any) *PluginResponse { _ = "STUB: not implemented"; return nil }

func NewStreamResponse(value any) *PluginResponse { _ = "STUB: not implemented"; return nil }

func (r *PluginResponse) Attr(name string) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// Error value is being checked in the handler code, clear the thread local state

// Value is being accessed when there was an error, abort

func (r *PluginResponse) AttrNames() []string { _ = "STUB: not implemented"; return nil }

func (r *PluginResponse) String() string { _ = "STUB: not implemented"; return "" }

func (r *PluginResponse) Type() string { _ = "STUB: not implemented"; return "" }

func (r *PluginResponse) Freeze() { _ = "STUB: not implemented"; return }

func (r *PluginResponse) Truth() starlark.Bool {
	_ = "STUB: not implemented"
	// Error value is being checked in the handler code, clear the thread local state
	return *new(starlark.Bool)
}

func (r *PluginResponse) Hash() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *PluginResponse) UnmarshalStarlarkType() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

var _ starlark.Value = (*PluginResponse)(nil)
var _ starlark_type.TypeUnmarshaler = (*PluginResponse)(nil)
