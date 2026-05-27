// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"net/http"
	"time"

	"github.com/openrundev/openrun/internal/types"
	"github.com/segmentio/ksuid"
)

var ridPrefix string

func init() {
	id, err := ksuid.NewRandom()
	if err != nil {
		panic(err)
	}
	ridPrefix = "rid_" + id.String() + "_"
}

func (s *Server) initAuditDB(connectString string) error { _ = "STUB: not implemented"; return nil }

const CURRENT_AUDIT_DB_VERSION = 1

func (s *Server) versionUpgradeAuditDB() error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck // ignore error if no version is found

//nolint:errcheck

func (s *Server) InsertAuditEvent(event *types.AuditEvent) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) cleanupEvents() error { _ = "STUB: not implemented"; return nil }

func (s *Server) auditCleanupLoop(cleanupTicker *time.Ticker) { _ = "STUB: not implemented"; return }

type ContextShared struct {
	UserId    string
	AppId     string
	Operation string
	Target    string
	DryRun    bool
}

func updateTargetInContext(r *http.Request, target string, dryRun bool) {
	_ = "STUB: not implemented"
	return
}

func updateOperationInContext(r *http.Request, operation string) { _ = "STUB: not implemented"; return }

var requestCounter uint64

func (server *Server) handleStatus(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Add a request id to the context

// Wrap the ResponseWriter

// Call the next handler

// Don't create audit events for get requests

// http event auditing is disabled for this app
