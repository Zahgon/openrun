// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"context"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/openrundev/openrun/internal/types"
)

func NewCertStorage(ctx context.Context, logger *types.Logger, metadata *Metadata) (*CertStorage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CertStorage is the database backed storage for the cert info. Implements the certmagic.Storage interface.
type CertStorage struct {
	*types.Logger
	metadata     *Metadata
	queryTimeout time.Duration
	lockTimeout  time.Duration
}

var _ certmagic.Storage = (*CertStorage)(nil)

// createTables creates the tables for the certificates.
func (s *CertStorage) createTables(ctx context.Context, tx types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// Lock locks the certificate with the given id.
func (c *CertStorage) Lock(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// Unlock unlocks the certificate with the given id.
func (c *CertStorage) Unlock(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// isLocked checks if the certificate with the given id is locked.
func (c *CertStorage) isLocked(ctx context.Context, tx types.Transaction, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Store stores the certificate with the given id and value.
func (c *CertStorage) Store(ctx context.Context, id string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Load loads the certificate with the given id.
func (c *CertStorage) Load(ctx context.Context, id string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// certmagic expects fs.ErrNotExist for missing keys

// Delete deletes the certificate with the given id.
func (c *CertStorage) Delete(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// Exists checks if the certificate with the given id exists.
func (c *CertStorage) Exists(ctx context.Context, id string) bool {
	_ = "STUB: not implemented"
	return false
}

// List lists the certificates with the given prefix.
func (c *CertStorage) List(ctx context.Context, prefix string, recursive bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Stat returns the information about the certificate with the given id.
func (c *CertStorage) Stat(ctx context.Context, id string) (certmagic.KeyInfo, error) {
	_ = "STUB: not implemented"
	return *new(certmagic.KeyInfo), nil
}
