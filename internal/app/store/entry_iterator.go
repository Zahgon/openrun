// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"database/sql"

	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

type StoreEntryIterable struct {
	thread *starlark.Thread
	*types.Logger
	table string
	rows  *sql.Rows
}

func NewStoreEntryIterabe(thread *starlark.Thread, logger *types.Logger, table string, rows *sql.Rows) *StoreEntryIterable {
	_ = "STUB: not implemented"
	return nil
}

var _ starlark.Iterable = (*StoreEntryIterable)(nil)

func (s *StoreEntryIterable) Iterate() starlark.Iterator {
	_ = "STUB: not implemented"
	return *new(starlark.Iterator)
}

func (s *StoreEntryIterable) String() string { _ = "STUB: not implemented"; return "" }

func (s *StoreEntryIterable) Type() string { _ = "STUB: not implemented"; return "" }

func (s *StoreEntryIterable) Freeze() {
	_ = "STUB: not implemented"
	// Not supported
	return
}

func (s *StoreEntryIterable) Truth() starlark.Bool {
	_ = "STUB: not implemented"
	return *new(starlark.Bool)
}

func (s *StoreEntryIterable) Hash() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

type StoreEntryIterator struct {
	thread *starlark.Thread
	*types.Logger
	table string
	rows  *sql.Rows
}

var _ starlark.Iterator = (*StoreEntryIterator)(nil)

func NewStoreEntryIterator(thread *starlark.Thread, logger *types.Logger, table string, rows *sql.Rows) *StoreEntryIterator {
	_ = "STUB: not implemented"
	return nil
}

func (i *StoreEntryIterator) Next(value *starlark.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func (i *StoreEntryIterator) Done() {
	_ = "STUB: not implemented"
	// Clear the deferred cleanup function, since Close is called here
	return
}
