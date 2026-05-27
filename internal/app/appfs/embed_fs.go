// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package appfs

import (
	"embed"
	"io/fs"

	"github.com/openrundev/openrun/internal/types"
)

type EmbedReadFS struct {
	*types.Logger
	fs embed.FS
}

var _ ReadableFS = (*EmbedReadFS)(nil)

func NewEmbedReadFS(logger *types.Logger, embedFS embed.FS) *EmbedReadFS {
	_ = "STUB: not implemented"
	return nil
}

func (e *EmbedReadFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

func (e *EmbedReadFS) ReadFile(name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (e *EmbedReadFS) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (e *EmbedReadFS) StatNoSpec(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (e *EmbedReadFS) Glob(pattern string) (matches []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EmbedReadFS) ReadDir(name string) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EmbedReadFS) StaticFiles() []string { _ = "STUB: not implemented"; return nil }

func (e *EmbedReadFS) FileHash(excludeGlob []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *EmbedReadFS) CreateTempSourceDir() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *EmbedReadFS) Reset() {
	_ = "STUB: not implemented"
	// do nothing
	return
}
