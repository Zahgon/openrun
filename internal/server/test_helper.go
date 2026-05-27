// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"time"
)

type InmemoryKVStore struct {
	store    map[string][]byte
	deleteAt map[string]*time.Time
}

func NewInmemoryKVStore() *InmemoryKVStore { _ = "STUB: not implemented"; return nil }

var _ KVStore = (*InmemoryKVStore)(nil)

func (s *InmemoryKVStore) FetchKV(ctx context.Context, key string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *InmemoryKVStore) FetchKVBlob(ctx context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *InmemoryKVStore) StoreKV(ctx context.Context, key string, value map[string]any, expireAt *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InmemoryKVStore) StoreKVBlob(ctx context.Context, key string, value []byte, expireAt *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InmemoryKVStore) UpsertKVBlob(ctx context.Context, key string, value []byte, expireAt *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InmemoryKVStore) UpdateKV(ctx context.Context, key string, value map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InmemoryKVStore) UpdateKVBlob(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InmemoryKVStore) DeleteKV(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InmemoryKVStore) ensureMaps() { _ = "STUB: not implemented"; return }

func copyTime(t *time.Time) *time.Time { _ = "STUB: not implemented"; return nil }
