// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
)

func (s *Server) ReloadApp(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, stageAppEntry *types.AppEntry,
	approve, dryRun, promote bool, branch, commit, gitAuth string, repoCache *RepoCache, forceReload bool) (*types.AppReloadResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Persist the metadata so that any git info is saved

// Persist name in metadata

// Persist name in metadata

func (s *Server) ReloadApps(ctx context.Context, appPathGlob string, approve, dryRun, promote bool,
	branch, commit, gitAuth string, forceReload bool) (*types.AppReloadResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Track the staging and prod apps

// Commit the transaction if not dry run and update the in memory app store

func (s *Server) loadAppCode(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, branch, commit, gitAuth string, repoCache *RepoCache, forceReload bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Commit is specified and matches the current version, skip reload

// If no commit is specified, and the current version is the same as the latest commit, skip reload

// Checkout the git repo locally and load into database

// App is loaded from disk (not git), load files into DB

func (s *Server) StagedUpdate(ctx context.Context, appPathGlob string, dryRun, promote bool, handler stagedUpdateHandler, args map[string]any, op string) (*types.AppStagedUpdateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

type stagedUpdateHandler func(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, args map[string]any) (any, types.AppPathDomain, error)

func (s *Server) StagedUpdateAppsTx(ctx context.Context, tx types.Transaction, appPathGlob string, promote bool, handler stagedUpdateHandler, args map[string]any) ([]any, []types.AppPathDomain, []types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// For prod apps, update the staging app

// prod app audit result is not added to results, since it will be same as the staging app

func (s *Server) auditHandler(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, args map[string]any) (any, types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return *new(any), *new(types.AppPathDomain), nil
}

func (s *Server) PromoteApps(ctx context.Context, appPathGlob string, dryRun bool) (*types.AppPromoteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Not a prod app, skip

func (s *Server) promoteApp(ctx context.Context, tx types.Transaction, stagingApp *types.AppEntry, prodApp *types.AppEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// the prod app version after promote is the same as the staging app version
// there might be some gaps in the prod app version numbers, but that is ok, the attempt is to have the version number in
// sync with the staging app version number when a promote is done

// Even if there is no version change, promotion is done to update other metadata settings like account links

func (s *Server) UpdateAppSettings(ctx context.Context, appPathGlob string, dryRun bool, updateAppRequest types.UpdateAppRequest) (*types.AppUpdateSettingsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Delete instead of update to avoid having to initialize all the linked apps

// Apps will get reloaded on the next request

func (s *Server) updateAppSettings(ctx context.Context, tx types.Transaction, appPathDomain types.AppPathDomain, updateAppRequest types.UpdateAppRequest) ([]types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Include the main app

func (s *Server) accountLinkHandler(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, args map[string]any) (any, types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return *new(any), *new(types.AppPathDomain), nil
}

// Update existing value

// Add new value

// Delete the entry

// Update existing value

func (s *Server) updateParamHandler(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, args map[string]any) (any, types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return *new(any), *new(types.AppPathDomain), nil
}

// Delete the entry

// Update existing value

func (s *Server) updateMetadataHandler(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, args map[string]any) (any, types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return *new(any), *new(types.AppPathDomain), nil
}

// The type is being updated

// updateAppMetadataConfig updates the app metadata config.
func (s *Server) updateAppMetadataConfig(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, configType types.AppMetadataConfigType, configEntries []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) validateAppBindings(ctx context.Context, tx types.Transaction, bindingPaths []string) error {
	_ = "STUB: not implemented"
	return nil
}
