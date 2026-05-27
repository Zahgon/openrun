// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"sync"

	"github.com/openrundev/openrun/internal/app"
	"github.com/openrundev/openrun/internal/container"
	"github.com/openrundev/openrun/internal/types"
)

// AppStore is a store of apps. List of apps is stored in memory. Apps are initialized lazily,
// AddApp has to be called before GetApp to initialize the app
type AppStore struct {
	*types.Logger
	server     *Server
	allApps    []types.AppInfo
	idToInfo   map[types.AppId]types.AppInfo
	allDomains map[string]bool

	mu     sync.RWMutex
	appMap map[types.AppPathDomain]*app.App
}

func NewAppStore(logger *types.Logger, server *Server) *AppStore {
	_ = "STUB: not implemented"
	return nil
}

func (a *AppStore) GetAppsFullInfo() ([]types.AppInfo, map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get exclusive lock

func (a *AppStore) GetAllAppsInfo() ([]types.AppInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get exclusive lock

func (a *AppStore) GetAppInfo(appId types.AppId) (types.AppInfo, bool) {
	_ = "STUB: not implemented"
	return *new(types.AppInfo), false
}

// Get exclusive lock

func (a *AppStore) GetAllDomains() (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get exclusive lock

func (a *AppStore) reloadAppInfo() error { _ = "STUB: not implemented"; return nil }

func (a *AppStore) ResetAllAppCache() { _ = "STUB: not implemented"; return }

func (a *AppStore) resetAllAppCache() { _ = "STUB: not implemented"; return }

func (a *AppStore) GetApp(pathDomain types.AppPathDomain) (*app.App, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ActiveContainerNames returns the container names currently referenced by loaded apps.
func (a *AppStore) ActiveContainerNames() map[container.ContainerName]bool {
	_ = "STUB: not implemented"
	return nil
}

func (a *AppStore) AddApp(app *app.App) { _ = "STUB: not implemented"; return }

func (a *AppStore) ClearLinkedApps(pathDomain types.AppPathDomain) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: create audit entry for linked apps

func (a *AppStore) clearApp(pathDomain types.AppPathDomain) { _ = "STUB: not implemented"; return }

//nolint:errcheck

// ClearApps removes the specified apps from the in memory App cache
// Also clears the app info cache for all apps (so that it is reloaded on next request)
func (a *AppStore) ClearApps(pathDomains []types.AppPathDomain) { _ = "STUB: not implemented"; return }

// ClearApps removes the specified apps from the in memory App cache
// Also clears the app info cache for all apps (so that it is reloaded on next request)
// This does not notify other servers of the app update (intended for use from the listener)
func (a *AppStore) ClearAppsNoNotify(pathDomains []types.AppPathDomain) {
	_ = "STUB: not implemented"
	return
}

// ClearApps removes the specified apps from the in memory App cache and creates an audit entry.
// Also clears the app info cache for all apps (so that it is reloaded on next request)
func (a *AppStore) ClearAppsAudit(ctx context.Context, pathDomains []types.AppPathDomain, op string) error {
	_ = "STUB: not implemented"
	return nil
}

func getAppInfoMap(appInfo []types.AppInfo) map[string]types.AppInfo {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck
