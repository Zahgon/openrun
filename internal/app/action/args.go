// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package action

import (
	"go.starlark.net/starlark"
)

// Args is a starlark.Value that represents the arguments being passed to the Action handler. It contains value for the params.
type Args struct {
	members starlark.StringDict
}

func (a *Args) Attr(name string) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (a *Args) AttrNames() []string { _ = "STUB: not implemented"; return nil }

func (a *Args) String() string { _ = "STUB: not implemented"; return "" }

func (a *Args) Type() string { _ = "STUB: not implemented"; return "" }

func (a *Args) Freeze() { _ = "STUB: not implemented"; return }

func (a *Args) Truth() starlark.Bool { _ = "STUB: not implemented"; return *new(starlark.Bool) }

func (a *Args) Hash() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

var _ starlark.Value = (*Args)(nil)
