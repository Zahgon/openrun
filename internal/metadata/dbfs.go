// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"bytes"
	"io"
	"io/fs"
	"time"

	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/types"
)

type DbFs struct {
	*types.Logger
	fileStore *FileStore
	fileInfo  map[string]DbFileInfo
	specFiles types.SpecFiles
}

var _ appfs.ReadableFS = (*DbFs)(nil)

func NewDbFs(logger *types.Logger, fileStore *FileStore, specFiles types.SpecFiles) (*DbFs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DbFile struct {
	name   string
	fi     DbFileInfo
	reader *DbFileReader
}

var _ fs.File = (*DbFile)(nil)

func NewDBFile(name string, compressionType string, data []byte, fi DbFileInfo) *DbFile {
	_ = "STUB: not implemented"
	return nil
}

func (f *DbFile) Read(dst []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *DbFile) Name() string { _ = "STUB: not implemented"; return "" }

func (f *DbFile) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (f *DbFile) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	// Seek is called by http.ServeContent in source_fs for the unoptimized case only
	// The data is decompressed and then recompressed if required in the unoptimized case
	return 0, nil
}

func (f *DbFile) ReadCompressed() ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (f *DbFile) Close() error { _ = "STUB: not implemented"; return nil }

type DbFileReader struct {
	compressionType    string
	compressedReader   *bytes.Reader
	uncompressedReader *bytes.Reader
}

var _ io.ReadSeeker = (*DbFileReader)(nil)
var _ appfs.CompressedReader = (*DbFileReader)(nil)

func NewDbFileReader(compressionType string, data []byte) *DbFileReader {
	_ = "STUB: not implemented"
	return nil
}

func (f *DbFileReader) uncompress() error { _ = "STUB: not implemented"; return nil }

func (f *DbFileReader) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *DbFileReader) Read(dst []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *DbFileReader) ReadCompressed() ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

type DbFileInfo struct {
	name    string
	len     int64
	sha     string
	modTime time.Time
}

var _ fs.FileInfo = (*DbFileInfo)(nil)

func (fi *DbFileInfo) Name() string { _ = "STUB: not implemented"; return "" }

func (fi *DbFileInfo) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (fi *DbFileInfo) Mode() fs.FileMode { _ = "STUB: not implemented"; return *new(fs.FileMode) }

func (fi *DbFileInfo) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (fi *DbFileInfo) IsDir() bool { _ = "STUB: not implemented"; return false }

func (fi *DbFileInfo) Sys() any { _ = "STUB: not implemented"; return *new(any) }

func computeSha(data string) string { _ = "STUB: not implemented"; return "" }

func (d *DbFs) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

// Check if the spec files has it

func (d *DbFs) ReadFile(name string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Check if the spec files has it

func (d *DbFs) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

// Check if the spec files has it

func (d *DbFs) StatNoSpec(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (d *DbFs) Glob(pattern string) (matches []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DbFs) StaticFiles() []string { _ = "STUB: not implemented"; return nil }

// GlobMatch returns true if the file name matches any of the patterns
func GlobMatch(patterns []string, fileName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// FileHash returns a hash of the file names and their corresponding sha256 hashes
func (d *DbFs) FileHash(excludeGlob []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Name is excluded from the hash, must be a file used by the openrun hypermedia based UI
// We don't want a UI only change to cause a container rebuild

// Name is excluded from the hash, must be a file used by the openrun hypermedia based UI
// We don't want a UI only change to cause a container rebuild

// Only include spec files that are not already in the file info

func (d *DbFs) CreateTempSourceDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Skip files that are already in the file info

func (d *DbFs) Reset() { _ = "STUB: not implemented"; return }
