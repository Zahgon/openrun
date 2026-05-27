// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"context"
	"database/sql"

	"github.com/openrundev/openrun/internal/types"
)

const (
	COMPRESSION_THRESHOLD = 0 // files above this size are stored compressed. The chi Compress middleware does not have a threshold,
	// so the threshold is set to zero here
	BROTLI_COMPRESSION_LEVEL = 9 // https://paulcalvano.com/2018-07-25-brotli-compression-how-much-will-it-reduce-your-content/ seems
	// to indicate that level 9 is a good default.

	defaultUser = "admin"
)

type FileStore struct {
	appId    types.AppId
	version  int
	metadata *Metadata
	db       *sql.DB
	initTx   types.Transaction // This is the transaction for the initial setup of the app, before it is committed to the database.
	// After app is committed to database, this is not used, auto-commit transactions are used for reads
	fileCache *FileCache
}

func NewFileStore(appId types.AppId, version int, metadata *Metadata, tx types.Transaction) (*FileStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FileStore) IncrementAppVersion(ctx context.Context, tx types.Transaction, metadata *types.AppMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

type fileEntry struct {
	path           string
	sha            string
	uncompressedSz int
	compression    string
	compressed     []byte
	shaExists      bool
	err            error
}

func (f *FileStore) AddAppVersionDisk(ctx context.Context, tx types.Transaction, metadata types.AppMetadata, checkoutDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect all file paths first

// Build set of existing SHAs in the files table

//nolint:errcheck

// Prepare insert statements upfront

//nolint:errcheck

//nolint:errcheck

// done is closed on early return to unblock goroutines and prevent leaks.

// work feeds paths to workers, results is bounded to numWorkers
// so at most numWorkers compressed files are held in memory at once.

//nolint:errcheck

// Consume results and insert into DB as they arrive

func (f *FileStore) GetFileBySha(sha string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

//nolint:errcheck

func (f *FileStore) GetFileByShaTx(ctx context.Context, tx types.Transaction, sha string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

//nolint:errcheck

func (f *FileStore) getFileInfo() (map[string]DbFileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (f *FileStore) getFileInfoTx(ctx context.Context, tx types.Transaction) (map[string]DbFileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

func (f *FileStore) GetHighestVersion(ctx context.Context, tx types.Transaction, appId types.AppId) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// No versions found

func (f *FileStore) PromoteApp(ctx context.Context, tx types.Transaction, prodAppId types.AppId, metadata *types.AppMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

// Use direct queries instead of prepared statements to avoid connection issues

//nolint:errcheck

// Collect all file data first to avoid keeping the result set open

// Now insert all files

func (f *FileStore) GetAppVersions(ctx context.Context, tx types.Transaction) ([]types.AppVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (f *FileStore) GetAppVersion(ctx context.Context, tx types.Transaction, version int) (*types.AppVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FileStore) GetAppFiles(ctx context.Context, tx types.Transaction) ([]types.AppFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FileStore) Reset() {
	_ = "STUB: not implemented"
	// Unlink the file store from the types.Transaction used during init
	return
}
