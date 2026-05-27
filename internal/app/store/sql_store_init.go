// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"context"

	"github.com/openrundev/openrun/internal/app/starlark_type"
)

func (s *SqlStore) initStore(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *SqlStore) createSchemaInfo(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Schema is up to date. This means there is an existing entry and that has a has same as the current schema

// Either no existing schema entry or hash mismatch. Insert new entry

func createIndexStmt(unquotedTableName string, index starlark_type.Index, mapper fieldMapper, quoteIdentifier func(string) string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
