// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

const (
	DEFAULT_FILE_LIMIT = 10_000
	MAX_FILE_LIMIT     = 100_000
)

type AccessType string

const (
	UserAccess AccessType = "user"
	AppAccess  AccessType = "app"
)

func initFS() { _ = "STUB: not implemented"; return }

type fsPlugin struct {
	accessAllowed []string
	pluginContext *types.PluginContext
}

func NewFSPlugin(pluginContext *types.PluginContext) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func resolveDirs(allowed []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Resolve symbolic links and canonicalize the paths

func (f *fsPlugin) checkAccess(filePath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *fsPlugin) Abs(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (f *fsPlugin) List(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (f *fsPlugin) Find(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

type FileInfo struct {
	Name  string
	Size  int64
	IsDir bool
	Mode  int
}

func listDir(ctx context.Context, path string, recursiveSize, ignoreError bool) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// syscall.Statfs is not available on Windows, using 4K as block size

func dirSize(ctx context.Context, path string, blockSize int64, ignoreError bool) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func convertToBlockSize(size, blockSize int64) int64 { _ = "STUB: not implemented"; return 0 }

func find(ctx context.Context, path, nameGlob string, limit, minSize int64, ignoreError bool) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// syscall.Statfs is not available on Windows, using 4K as block size

func truncateList(entries []*FileInfo, limit int64) []*FileInfo {
	_ = "STUB: not implemented"
	return nil
}

func matchFile(name, nameGlob string, size, minSize int64) bool {
	_ = "STUB: not implemented"
	return false
}

func matchFiles(ctx context.Context, path string, nameGlob string, limit, minSize int64, ignoreError bool) ([]*FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
