// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/openrundev/openrun/internal/app/starlark_type"
	"go.starlark.net/starlark"
)

const (
	ID_FIELD         = "_id"
	VERSION_FIELD    = "_version"
	CREATED_BY_FIELD = "_created_by"
	UPDATED_BY_FIELD = "_updated_by"
	CREATED_AT_FIELD = "_created_at"
	UPDATED_AT_FIELD = "_updated_at"
	JSON_FIELD       = "_json"
)

var RESERVED_FIELDS = map[string]bool{
	ID_FIELD:         true,
	VERSION_FIELD:    true,
	CREATED_BY_FIELD: true,
	UPDATED_BY_FIELD: true,
	CREATED_AT_FIELD: true,
	UPDATED_AT_FIELD: true,
	JSON_FIELD:       true,
}

type EntryId int64
type UserId string
type Document map[string]any

type Entry struct {
	Id        EntryId
	Version   int64
	CreatedBy UserId
	UpdatedBy UserId
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      Document
}

var _ starlark.Unpacker = (*Entry)(nil)

func (e *Entry) Unpack(value starlark.Value) error { _ = "STUB: not implemented"; return nil }

// Store is the interface for a OpenRun document store. These API are exposed by the db plugin
type Store interface {
	// Begin starts a new transaction
	Begin(ctx context.Context) (*sql.Tx, error)

	// Commit commits a transaction
	Commit(ctx context.Context, tx *sql.Tx) error

	// Rollback rolls back a transaction
	Rollback(ctx context.Context, tx *sql.Tx) error

	// Insert a new entry in the store
	Insert(ctx context.Context, tx *sql.Tx, table string, Entry *Entry) (EntryId, error)

	// SelectById returns a single item from the store
	SelectById(ctx context.Context, tx *sql.Tx, table string, id EntryId) (*Entry, error)

	// SelectOne returns a single item from the store
	SelectOne(ctx context.Context, tx *sql.Tx, table string, filter map[string]any) (*Entry, error)

	// Select returns the entries matching the filter
	Select(ctx context.Context, tx *sql.Tx, thread *starlark.Thread, table string, filter map[string]any, sort []string, offset, limit int64) (starlark.Iterable, error)

	// Count returns the count of entries matching the filter
	Count(ctx context.Context, tx *sql.Tx, table string, filter map[string]any) (int64, error)

	// Update an existing entry in the store
	Update(ctx context.Context, tx *sql.Tx, table string, Entry *Entry) (int64, error)

	// DeleteById an entry from the store by id
	DeleteById(ctx context.Context, tx *sql.Tx, table string, id EntryId) (int64, error)

	// Delete entries from the store matching the filter
	Delete(ctx context.Context, tx *sql.Tx, table string, filter map[string]any) (int64, error)
}

func CreateType(name string, entry *Entry) (*starlark_type.StarlarkType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO - add missing fields
