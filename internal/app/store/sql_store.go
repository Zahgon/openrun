// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"context"
	"database/sql"
	"sync"

	"github.com/openrundev/openrun/internal/system"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

const (
	SELECT_MAX_LIMIT     = 100_000
	SELECT_DEFAULT_LIMIT = 10_000
	SORT_ASCENDING       = "asc"
	SORT_DESCENDING      = "desc"
	MAX_TABLE_NAME_LEN   = 63
)

type SqlStore struct {
	*types.Logger
	sync.Mutex
	isInitialized bool
	pluginContext *types.PluginContext
	db            *sql.DB
	prefix        string
	isSqlite      bool // false means postgres, no other options
}

var _ Store = (*SqlStore)(nil)

func NewSqlStore(pluginContext *types.PluginContext) (*SqlStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SqlStore) dbType() system.DBType { _ = "STUB: not implemented"; return *new(system.DBType) }

func (s *SqlStore) rebindQuery(query string) string { _ = "STUB: not implemented"; return "" }

func (s *SqlStore) queryMapper() fieldMapper { _ = "STUB: not implemented"; return *new(fieldMapper) }

func (s *SqlStore) sortMapper() fieldMapper { _ = "STUB: not implemented"; return *new(fieldMapper) }

func (s *SqlStore) queryOptions() queryOptions {
	_ = "STUB: not implemented"
	return *new(queryOptions)
}

func (s *SqlStore) quoteIdentifier(identifier string) string { _ = "STUB: not implemented"; return "" }

func (s *SqlStore) genRawTableName(table string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateTableName(name string) error { _ = "STUB: not implemented"; return nil }

func isTableNameStart(ch byte) bool { _ = "STUB: not implemented"; return false }

func isTableNameChar(ch byte) bool { _ = "STUB: not implemented"; return false }

func genSortString(sortFields []string, mapper fieldMapper) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// :ASC is optional

func (s *SqlStore) genTableName(table string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *SqlStore) initialize(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Already initialized

func (s *SqlStore) Begin(ctx context.Context) (*sql.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SqlStore) Commit(ctx context.Context, tx *sql.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SqlStore) Rollback(ctx context.Context, tx *sql.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

// Insert a new entry in the store
func (s *SqlStore) Insert(ctx context.Context, tx *sql.Tx, table string, entry *Entry) (EntryId, error) {
	_ = "STUB: not implemented"
	return *new(EntryId), nil
}

// TODO update userid

// SelectById returns a single item from the store
func (s *SqlStore) SelectById(ctx context.Context, tx *sql.Tx, table string, id EntryId) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SelectOne returns a single item from the store
func (s *SqlStore) SelectOne(ctx context.Context, tx *sql.Tx, table string, filter map[string]any) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Select returns the entries matching the filter
func (s *SqlStore) Select(ctx context.Context, tx *sql.Tx, thread *starlark.Thread, table string, filter map[string]any, sort []string, offset, limit int64) (starlark.Iterable, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Iterable), nil
}

// Count returns the number of entries matching the filter
func (s *SqlStore) Count(ctx context.Context, tx *sql.Tx, table string, filter map[string]any) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Update an existing entry in the store
func (s *SqlStore) Update(ctx context.Context, tx *sql.Tx, table string, entry *Entry) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO update userid

// DeleteById an entry from the store by id
func (s *SqlStore) DeleteById(ctx context.Context, tx *sql.Tx, table string, id EntryId) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Delete entries from the store matching the filter
func (s *SqlStore) Delete(ctx context.Context, tx *sql.Tx, table string, filter map[string]any) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
