// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"encoding/base32"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const (
	kvSessionIDBytes        = 32
	maxKVSessionValueLength = 10 << 20 // 10MB
)

var kvSessionIDEncoding = base32.StdEncoding.WithPadding(base32.NoPadding)

// KVSessionStore keeps session payloads server-side and puts only an opaque,
// signed session id in the browser cookie.
type KVSessionStore struct {
	Codecs      []securecookie.Codec
	valueCodecs []securecookie.Codec
	Options     *sessions.Options
	db          KVStore
}

var _ sessions.Store = (*KVSessionStore)(nil)

func NewKVSessionStore(db KVStore, keyPairs ...[]byte) *KVSessionStore {
	_ = "STUB: not implemented"
	return nil
}

// Server-side values can exceed cookie limits, but should still be bounded.

func (s *KVSessionStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *KVSessionStore) New(r *http.Request, name string) (*sessions.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *KVSessionStore) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *KVSessionStore) MaxAge(age int) { _ = "STUB: not implemented"; return }

func (s *KVSessionStore) load(r *http.Request, session *sessions.Session) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *KVSessionStore) save(r *http.Request, session *sessions.Session) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *KVSessionStore) kvKey(session *sessions.Session) string {
	_ = "STUB: not implemented"
	return ""
}
