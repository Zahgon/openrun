// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"context"
	"database/sql"

	"github.com/openrundev/openrun/internal/types"
)

const CURRENT_FILE_CACHE_VERSION = 1

// TODO : add file cleanup logic

type FileCache struct {
	db *sql.DB
	*types.Logger
}

func InitFileCache(logger *types.Logger, config *types.ServerConfig) (*FileCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FileCache) VersionUpgrade(config *types.ServerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck // ignore error if no version is found

//nolint:errcheck

func (f *FileCache) GetCachedFile(ctx context.Context, sha string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

//nolint:errcheck

func (f *FileCache) AddCache(ctx context.Context, sha string, compressionType string, content []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck
