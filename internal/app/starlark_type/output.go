// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package starlark_type

import (
	"go.starlark.net/starlark"
)

type Output struct {
	Value starlark.Value
	Err   string
}

func (o Output) Attr(name string) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (o Output) AttrNames() []string { _ = "STUB: not implemented"; return nil }

func (o Output) String() string { _ = "STUB: not implemented"; return "" }

func (o Output) Type() string { _ = "STUB: not implemented"; return "" }

func (o Output) Freeze() { _ = "STUB: not implemented"; return }

func (o Output) Truth() starlark.Bool { _ = "STUB: not implemented"; return *new(starlark.Bool) }

func (o Output) Hash() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

var _ starlark.Value = (*Output)(nil)
