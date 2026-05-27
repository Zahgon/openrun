// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package appfs

import (
	"bytes"
	"io/fs"
	"time"

	"github.com/openrundev/openrun/internal/types"
)

type DiskReadFS struct {
	*types.Logger
	root      string
	specFiles types.SpecFiles
}

var _ ReadableFS = (*DiskReadFS)(nil)

func NewDiskReadFS(logger *types.Logger, root string, specFiles types.SpecFiles) *DiskReadFS {
	_ = "STUB: not implemented"
	return nil
}

type DiskWriteFS struct {
	*DiskReadFS
}

func (d *DiskReadFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

//nolint:errcheck

// File found in spec files, use that

func (d *DiskReadFS) ReadFile(name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// File found in spec files, use that

func (d *DiskReadFS) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

//nolint:errcheck

func (d *DiskReadFS) StatNoSpec(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

//nolint:errcheck

func (d *DiskReadFS) Glob(pattern string) (matches []string, err error) {
	_ = "STUB: not implemented"
	// TODO glob does not look at spec files
	return nil, nil
}

//nolint:errcheck

func (d *DiskReadFS) StaticFiles() []string { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (d *DiskReadFS) FileHash(excludeGlob []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *DiskReadFS) CreateTempSourceDir() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *DiskReadFS) Reset() {
	_ = "STUB: not implemented"
	// do nothing
	return
}

func (d *DiskWriteFS) Write(name string, bytes []byte) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (d *DiskWriteFS) Remove(name string) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (d *DiskReadFS) cleanName(name string) (localName string, specName string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type DiskFile struct {
	name   string
	fi     DiskFileInfo
	reader *bytes.Reader
}

var _ fs.File = (*DiskFile)(nil)

func NewDiskFile(name string, data []byte, fi DiskFileInfo) *DiskFile {
	_ = "STUB: not implemented"
	return nil
}

func (f *DiskFile) Read(dst []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *DiskFile) Name() string { _ = "STUB: not implemented"; return "" }

func (f *DiskFile) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (f *DiskFile) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	// Seek is called by http.ServeContent in source_fs for the unoptimized case only
	// The data is decompressed and then recompressed if required in the unoptimized case
	return 0, nil
}

func (f *DiskFile) Close() error { _ = "STUB: not implemented"; return nil }

type DiskFileInfo struct {
	name    string
	len     int64
	modTime time.Time
}

var _ fs.FileInfo = (*DiskFileInfo)(nil)

func (fi *DiskFileInfo) Name() string { _ = "STUB: not implemented"; return "" }

func (fi *DiskFileInfo) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (fi *DiskFileInfo) Mode() fs.FileMode { _ = "STUB: not implemented"; return *new(fs.FileMode) }

func (fi *DiskFileInfo) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (fi *DiskFileInfo) IsDir() bool { _ = "STUB: not implemented"; return false }

func (fi *DiskFileInfo) Sys() any { _ = "STUB: not implemented"; return *new(any) }
