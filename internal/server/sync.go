// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
)

func (s *Server) CreateSyncEntry(ctx context.Context, path string, scheduled, dryRun bool, sync *types.SyncMetadata) (*types.SyncCreateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Webhook sync entry

// Persist the settings

// The sync job job failed, delete the entry

// TODO

func (s *Server) RunSync(ctx context.Context, id string, dryRun bool) (*types.SyncJobStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// The sync job job failed, status would be already updated

func (s *Server) DeleteSyncEntry(ctx context.Context, id string, dryRun bool) (*types.SyncDeleteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) ListSyncEntries(ctx context.Context) (*types.SyncListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// TODO: Set the actual webhook URL

func (s *Server) syncRunner() { _ = "STUB: not implemented"; return }

func (s *Server) runSyncJobs() error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

// Create a new repo cache if not passed in

// each sync runs in its own transaction

// One failure does not stop the rest

func (s *Server) runSyncJob(ctx context.Context, inputTx types.Transaction, entry *types.SyncEntry,
	dryRun, checkCommitHash bool, repoCache *RepoCache) (*types.SyncJobStatus, []types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//nolint:errcheck

// No rollback here if transaction is passed in

// Create a new repo cache if not passed in

// This run was skipped, use the last run apps

// The apply was skipped, check if the apps need to be reloaded
// The attempt is to avoid doing a full github checkout on the apply file repo and on the
// app source repo, a list API is used to get the last commit

// App has been deleted, run the full apply with the latest commit even if it was already applied

// abort reloads

//nolint:errcheck // rollback any changes to db done during apply or reload
// CreateSyncEntry also aborts if the sync job fails, so rolling back the transaction here is fine
// Use a new transaction to update the sync status

//nolint:errcheck
