// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/openrundev/openrun/internal/system"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

var (
	mu       sync.RWMutex
	fsDB     *sql.DB
	fsDBType system.DBType
)

func InitFileStore(ctx context.Context, connectString string) error {
	_ = "STUB: not implemented"
	return nil
}

func fileCleanup(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func backgroundCleanup(ctx context.Context, cleanupTicker *time.Ticker) {
	_ = "STUB: not implemented"
	return
}

func (f *fsPlugin) ServeTmpFile(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func AddUserFile(ctx context.Context, file *types.UserFile) error {
	_ = "STUB: not implemented"
	return nil
}

func GetUserFile(ctx context.Context, id string) (*types.UserFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteUserFile(ctx context.Context, id string) error { _ = "STUB: not implemented"; return nil }

type expiredFile struct {
	Id       string
	FilePath string
}

func listExpiredFile(ctx context.Context) ([]expiredFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

func deleteExpiredFiles(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck
