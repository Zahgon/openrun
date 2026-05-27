// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlarkstruct"
)

const (
	APP = "app"
)

func (s *Server) loadApplyInfo(fileName string, data []byte, branch string, applyDev bool) ([]*types.CreateAppRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appDefToApplyInfo(appDef *starlarkstruct.Struct) (*types.CreateAppRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) setupSource(applyPath, branch, commit, gitAuth string, repoCache *RepoCache, isDev bool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (s *Server) Apply(ctx context.Context, inputTx types.Transaction, applyPath string, appPathGlob string, approve, dryRun, promote bool,
	reload types.AppReloadOption, branch, commit, gitAuth string, clobber,
	forceReload bool, lastRunCommitId string, repoCache *RepoCache, isDev bool) (*types.AppApplyResponse, []types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//nolint:errcheck

// No rollback here if transaction is passed in

// If no commit is specified, and the current version is the same as the latest commit, skip apply
// Only schedule sync passes in the lastRunCommitId, so this does not happen for normal apply

// If domain ends with a dot, append the default domain

// New app being created

// Override the dev status from the apply command cli

// Get list of all updated apps

// Commit the transaction if not dry run and update the in memory app store

func convertToMapString(input map[string]any, convertToml bool) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) applyAppUpdate(ctx context.Context, tx types.Transaction, appPathDomain types.AppPathDomain, newInfo *types.CreateAppRequest,
	approve, dryRun, promote bool, reload types.AppReloadOption, clobber bool, repoCache *RepoCache, forceReload bool) (*types.AppApplyResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For prod apps, update the staging app

// Reload does the version increment and promotion

// No reload, increment version and promote (if enabled)

func mergeMap(old, new, live map[string]string, clobber bool) bool {
	_ = "STUB: not implemented"

	// Force overwrite the live map
	return false
}

// Force update all values

// First run of apply

// Add values from new, retaining existing live values

// Three way merge

// Changed from old to new

// Removed from new

// Added in new

func mergeSlice(old, new []string, live *[]string, clobber bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Force update all values

// First run of apply

// Add values from new, retaining existing live values

// Three way merge

// Removed from new

// Added in new

func checkPropertyChanged(oldInfo *types.CreateAppRequest, fetchVal func(*types.CreateAppRequest) any, newVal, liveVal any, clobber bool) bool {
	_ = "STUB: not implemented"
	return false
}
