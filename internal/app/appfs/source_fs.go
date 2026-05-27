// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package appfs

import (
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"regexp"
	"sync"
	"time"
)

// ReadableFS is the interface for the file system used by app to read source files
type ReadableFS interface {
	fs.FS
	fs.ReadFileFS
	fs.GlobFS
	// Stat returns the stats for the named file.
	Stat(name string) (fs.FileInfo, error)
	StatNoSpec(name string) (fs.FileInfo, error)   // Stat the FS without looking at spec files
	Reset()                                        // Used to reset the file system transaction for the DbFs, no-op for others
	StaticFiles() []string                         // Return list of static files
	FileHash(excludeGlob []string) (string, error) // Return a hash of the source file contents
	CreateTempSourceDir() (string, error)          // Create a temporary directory with source files
}

type CompressedReader interface {
	ReadCompressed() (data []byte, compressionType string, err error)
}

// WritableFS is the interface for the writable underlying file system used by AppFS
type WritableFS interface {
	ReadableFS
	Write(name string, bytes []byte) error
	Remove(name string) error
}

// SourceFs is the implementation of source file system
type SourceFs struct {
	ReadableFS
	Root  string
	isDev bool

	staticFiles []string
	mu          sync.RWMutex
	nameToHash  map[string]string    // lookup (path to hash path)
	hashToName  map[string][2]string // reverse lookup (hash path to path)
}

var _ ReadableFS = (*SourceFs)(nil)

type WritableSourceFs struct {
	*SourceFs
}

var _ WritableFS = (*WritableSourceFs)(nil)

func (w *WritableSourceFs) Write(name string, bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *WritableSourceFs) Remove(name string) error { _ = "STUB: not implemented"; return nil }

func NewSourceFs(dir string, fs ReadableFS, isDev bool) (*SourceFs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// File hashing code based on https://github.com/benbjohnson/hashfs/blob/main/hashfs.go
// Copyright (c) 2020 Ben Johnson. MIT License

func (f *SourceFs) StaticFiles() []string { _ = "STUB: not implemented"; return nil }

func (f *SourceFs) ClearCache() { _ = "STUB: not implemented"; return }

func (f *SourceFs) Glob(pattern string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *SourceFs) ParseFS(funcMap template.FuncMap, patterns ...string) (*template.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *SourceFs) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

// Open returns a reference to the named file.
// If name is a hash name then the underlying file is used.
func (f *SourceFs) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

func (f *SourceFs) open(name string) (_ fs.File, hash string, err error) {
	_ = "STUB: not implemented"
	// Parse filename to see if it contains a hash.
	// If so, check if hash name matches.
	return *new(fs.File), "", nil
}

// HashName returns the hash name for a path, if exists.
// Otherwise returns the original path.
func (f *SourceFs) HashName(name string) string {
	_ = "STUB: not implemented"
	// Lookup cached formatted name, if exists.
	return ""
}

// Read file contents. Return original filename if we receive an error.

//TODO: log

// Compute hash and build filename.

// Store in lookups.

// FormatName returns a hash name that inserts hash before the filename's
// extension. If no extension exists on filename then the hash is appended.
// Returns blank string the original filename if hash is blank. Returns a blank
// string if the filename is blank.
func FormatName(filename, hash string) string { _ = "STUB: not implemented"; return "" }

// ParseName splits formatted hash filename into its base & hash components.
func (f *SourceFs) ParseName(filename string) (base, hash string) {
	_ = "STUB: not implemented"
	return "", ""
}

// ParseName splits formatted hash filename into its base & hash components.
func ParseName(filename string) (base, hash string) { _ = "STUB: not implemented"; return "", "" }

// Extract pre-hash & extension.

// If prehash doesn't contain the hash, then exit.

var hashSuffixRegex = regexp.MustCompile(`-[0-9a-f]{64}`)

// FileServer returns an http.Handler for serving FS files. It provides a
// simplified implementation of http.FileServer which is used to aggressively
// cache files on the client since the file hash is in the filename.
//
// Because FileServer is focused on small known path files, several features
// of http.FileServer have been removed including canonicalizing directories,
// defaulting index.html pages, precondition checks, & content range headers.
func FileServer(fsys *SourceFs, indexPage string) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func FileServerSingle(fsys *SourceFs, indexPage string) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type fsHandler struct {
	fsys       *SourceFs
	indexPage  string
	singleFile bool
}

func (h *fsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// Clean up filename based on URL path.
	return
}

// Read file from attached file system.

//nolint:errcheck

// Fetch file info. Disallow directories from being displayed.

// Cache the file aggressively if the file contains a hash.

// If this is a request without Range headers and brotli encoding is accepted,
// Return the data which is already in a compressed form

const COMPRESSION_TYPE = "br" // brotli uses br as the encoding type

func (h *fsHandler) canServeCompressed(r *http.Request) bool {
	_ = "STUB: not implemented"
	return false
}

// Range headers are being used, fallback to http.ServeContent

var unixEpochTime = time.Unix(0, 0)

// serveCompressed checks if the compressed file data can be streamed directly to the client, without
// the need to decompress and then recompress. If the client accepts brotli compressed data and there are no
// range headers, then this optimization can be used.
func (h *fsHandler) serveCompressed(w http.ResponseWriter, r *http.Request, filename string, modtime time.Time, content io.ReadSeeker) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// the data is not compressed with brotli, fallback to http.ServeContent

//nolint:errcheck
