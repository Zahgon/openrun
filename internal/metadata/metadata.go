// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package metadata

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/caddyserver/certmagic"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgxlisten"
	"github.com/openrundev/openrun/internal/system"
	"github.com/openrundev/openrun/internal/types"
	_ "modernc.org/sqlite"
)

const CURRENT_DB_VERSION = 11

// Metadata is the metadata persistence layer
type Metadata struct {
	*types.Logger
	certStorage      *CertStorage
	leaderElection   *LeaderElection
	config           *types.ServerConfig
	db               *sql.DB
	dbType           system.DBType
	pgListener       *pgxlisten.Listener
	AppNotifyFunc    func(types.AppUpdatePayload)
	ConfigNotifyFunc func(types.ConfigUpdatePayload)
}

const pg_listen_channel = "openrun_events"

// NewMetadata creates a new metadata persistence layer
func NewMetadata(logger *types.Logger, config *types.ServerConfig) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setup listener for app update notifications

// IsLeader returns true if the current server is the leader
func (m *Metadata) IsLeader() bool { _ = "STUB: not implemented"; return false }

// Close stops background goroutines owned by Metadata (e.g. leader election).
func (m *Metadata) Close() { _ = "STUB: not implemented"; return }

// GetCertStorage returns the cert storage implementation which persists the cert info to the database.
func (m *Metadata) GetCertStorage() certmagic.Storage {
	_ = "STUB: not implemented"
	return *

	// NotifyAppUpdate sends a notification through the postgres listener that an app has been updated
	new(certmagic.Storage)
}

func (m *Metadata) NotifyAppUpdate(appPathDomains []types.AppPathDomain) error {
	_ = "STUB: not implemented"
	return nil
}

// NotifyConfigUpdate sends a notification through the postgres listener that the config has been updated
func (m *Metadata) NotifyConfigUpdate() error { _ = "STUB: not implemented"; return nil }

func (m *Metadata) VersionUpgrade(config *types.ServerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck // ignore error if no version is found

//nolint:errcheck

func (m *Metadata) initFileTables(ctx context.Context, tx types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

type appMetadataAndSettings struct {
	path     string
	domain   string
	metadata *types.AppMetadata
	settings *types.AppSettings
}

func (m *Metadata) getAppMetadataAndSettings(ctx context.Context, tx types.Transaction) ([]appMetadataAndSettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

// migrateAuthSettings migrates the auth settings (app auth and git auth) from the app settings to the app metadata
func (m *Metadata) migrateAuthSettings(ctx context.Context, tx types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck // deprecated
//nolint:staticcheck // deprecated
// nolint:staticcheck // deprecated
//nolint:staticcheck // deprecated

func (m *Metadata) CreateApp(ctx context.Context, tx types.Transaction, app *types.AppEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) GetApp(pathDomain types.AppPathDomain) (*types.AppEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (m *Metadata) GetAppTx(ctx context.Context, tx types.Transaction, pathDomain types.AppPathDomain) (*types.AppEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (m *Metadata) DeleteApp(ctx context.Context, tx types.Transaction, id types.AppId) error {
	_ = "STUB: not implemented"
	return nil
}

// Clean up unused files. This can be done more aggressively, when older versions are deleted.
// Currently done only when an app is deleted. This cleanup is across apps, not just the deleted app.

func (m *Metadata) GetAppsForDomain(domain string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

func (m *Metadata) GetAllApps(includeInternal bool) ([]types.AppInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

// GetLinkedApps gets all the apps linked to the given main app (staging and preview apps)
func (m *Metadata) GetLinkedApps(ctx context.Context, tx types.Transaction, mainAppId types.AppId) ([]*types.AppEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

// No linked apps found, return empty slice

func (m *Metadata) UpdateSourceUrl(ctx context.Context, tx types.Transaction, app *types.AppEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) UpdateAppMetadata(ctx context.Context, tx types.Transaction, app *types.AppEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) updateAppMetadata(ctx context.Context, tx types.Transaction, path, domain string, metadata *types.AppMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) UpdateAppSettings(ctx context.Context, tx types.Transaction, app *types.AppEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) updateAppSettings(ctx context.Context, tx types.Transaction, path, domain string, settings *types.AppSettings) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) CreateSync(ctx context.Context, tx types.Transaction, sync *types.SyncEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) DeleteSync(ctx context.Context, tx types.Transaction, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSyncEntries gets all the sync entries for the given webhook type
func (m *Metadata) GetSyncEntries(ctx context.Context, tx types.Transaction) ([]*types.SyncEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

// No entries found, return empty slice

func (m *Metadata) GetSyncEntry(ctx context.Context, tx types.Transaction, id string) (*types.SyncEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Metadata) UpdateSyncStatus(ctx context.Context, tx types.Transaction, id string, status *types.SyncJobStatus) error {
	_ = "STUB: not implemented"
	return nil
}

var ErrConfigAlreadyExists = errors.New("config already exists")
var ErrConfigNotFound = errors.New("config not found")

func (m *Metadata) InitConfig(ctx context.Context, user string, dynamicConfig *types.DynamicConfig) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func (m *Metadata) UpdateConfig(ctx context.Context, user string, oldVersionId string, dynamicConfig *types.DynamicConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) GetConfig() (*types.DynamicConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Metadata) FetchKV(ctx context.Context, key string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Metadata) FetchKVBlob(ctx context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Metadata) StoreKV(ctx context.Context, key string, value map[string]any, expireAt *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) StoreKVBlob(ctx context.Context, key string, value []byte, expireAt *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) UpsertKVBlob(ctx context.Context, key string, value []byte, expireAt *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) UpdateKV(ctx context.Context, key string, value map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) UpdateKVBlob(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) DeleteKV(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) CleanupExpiredKV(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) CleanupAppVersions(app types.AppInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// The current version may not be the latest (e.g. after a rollback), so only
// consider versions <= the current one. Keep the current version plus
// RetainVersions older versions below it; delete everything older.
// OFFSET RetainVersions in a DESC-sorted list of versions <= current gives the
// oldest version to keep. If fewer versions exist, the subquery returns NULL
// and nothing is deleted.

func (m *Metadata) CleanupFiles() error { _ = "STUB: not implemented"; return nil }

func (m *Metadata) createServiceBindings(ctx context.Context, tx types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) CreateService(ctx context.Context, tx types.Transaction, service *types.Service) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) UpdateService(ctx context.Context, tx types.Transaction, service *types.Service) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) ServiceExists(ctx context.Context, tx types.Transaction, serviceType, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ClearServiceDefault unsets the is_default flag for any service of the given
// service_type except for the service with the given name. If exceptName is empty,
// the default flag is cleared for all services of that type.
func (m *Metadata) ClearServiceDefault(ctx context.Context, tx types.Transaction, serviceType, exceptName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) ClearServiceStaging(ctx context.Context, tx types.Transaction, serviceType, staging string) error {
	_ = "STUB: not implemented"
	return nil
}

// CountServices returns the number of services of the given service_type.
func (m *Metadata) CountServices(ctx context.Context, tx types.Transaction, serviceType string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Metadata) DeleteService(ctx context.Context, tx types.Transaction, name, serviceType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) GetDefaultService(ctx context.Context, tx types.Transaction, serviceType string) (*types.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Metadata) GetService(ctx context.Context, tx types.Transaction, serviceType, name string) (*types.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListServices returns services filtered by the optional serviceType and name. Empty string means no filter.
func (m *Metadata) ListServices(ctx context.Context, tx types.Transaction, serviceType, name string) ([]*types.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

func (m *Metadata) CreateBinding(ctx context.Context, tx types.Transaction, binding *types.Binding) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) UpdateBinding(ctx context.Context, tx types.Transaction, binding *types.Binding) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) DeleteBinding(ctx context.Context, tx types.Transaction, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Metadata) GetBinding(ctx context.Context, tx types.Transaction, path string) (*types.Binding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListBindings returns bindings filtered by the optional source. Empty string means no filter.
func (m *Metadata) ListBindings(ctx context.Context, tx types.Transaction, source string) ([]*types.Binding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

func toNullTime(t *time.Time) sql.NullTime { _ = "STUB: not implemented"; return *new(sql.NullTime) }

// BeginTransaction starts a new Transaction
func (m *Metadata) BeginTransaction(ctx context.Context) (types.Transaction, error) {
	_ = "STUB: not implemented"
	return *new(types.Transaction), nil
}

// CommitTransaction commits a transaction
func (m *Metadata) CommitTransaction(tx types.Transaction) error {
	_ = "STUB: not implemented"
	return nil

	// RollbackTransaction rolls back a transaction
}

func (m *Metadata) RollbackTransaction(tx types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}
