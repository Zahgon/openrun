// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"crypto/x509"
	"database/sql"
	"embed"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/openrundev/openrun/internal/app"
	"github.com/openrundev/openrun/internal/metadata"
	"github.com/openrundev/openrun/internal/rbac"
	"github.com/openrundev/openrun/internal/system"
	"github.com/openrundev/openrun/internal/telemetry"
	"github.com/openrundev/openrun/internal/types"
	"github.com/segmentio/ksuid"

	_ "github.com/openrundev/openrun/internal/app/store" // Register db plugin
	_ "github.com/openrundev/openrun/plugins"            // Register builtin plugins
)

const (
	DEFAULT_CERT_FILE = "default.crt"
	DEFAULT_KEY_FILE  = "default.key"
	APPSPECS          = "appspecs"
)

//go:embed appspecs
var embedAppTypes embed.FS

var appTypes map[string]types.SpecFiles

func init() {
	id, err := ksuid.NewRandom()
	if err != nil {
		panic(err)
	}
	types.CurrentServerId = types.ServerId(types.ID_PREFIX_SERVER + id.String())

	// Read app type config embedded in the binary
	appTypes = make(map[string]types.SpecFiles)
	entries, err := embedAppTypes.ReadDir(APPSPECS)
	if err != nil {
		return
	}

	for _, dir := range entries {
		// Loop through all directories in app specs, each is an app type
		if !dir.IsDir() || strings.HasPrefix(dir.Name(), ".") || dir.Name() == "dummy" {
			continue
		}
		files, err := embedAppTypes.ReadDir(path.Join(APPSPECS, dir.Name()))
		if err != nil {
			panic(err)
		}

		appType := make(types.SpecFiles)
		for _, file := range files {
			// Loop through all files in the app_type directory
			if file.IsDir() {
				continue
			}
			data, err := embedAppTypes.ReadFile(path.Join(APPSPECS, dir.Name(), file.Name()))
			if err != nil {
				panic(err)
			}
			appType[file.Name()] = string(data)
		}

		appTypes[dir.Name()] = appType
	}
}

func (s *Server) GetAppSpec(name types.AppSpec) types.SpecFiles {
	_ = "STUB: not implemented"
	// Add custom app type config from conf folder
	return *new(types.SpecFiles)
}

// Use bundled app if present

// Loop through all files in the app_type directory

// Server is the instance of the OpenRun Server
type Server struct {
	*types.Logger
	config         *types.ServerConfig
	db             *metadata.Metadata
	httpServer     *http.Server
	httpsServer    *http.Server
	udsServer      *http.Server
	handler        *Handler
	apps           *AppStore
	authHandler    *AdminBasicAuth
	oAuthManager   *OAuthManager
	samlManager    *SAMLManager
	notifyClose    chan types.AppPathDomain
	secretsManager *system.SecretManager
	listAppsApp    *app.App
	mu             sync.RWMutex
	auditDB        *sql.DB
	auditDbType    system.DBType
	syncTimer      *time.Ticker
	configMu       sync.RWMutex
	dynamicConfig  *types.DynamicConfig
	rbacManager    *rbac.RBACManager
	csrfMiddleware *http.CrossOriginProtection
	telemetry      *telemetry.Providers

	forwardAuthHTTPClient *http.Client

	staleContainerCleanupTicker *time.Ticker
	staleContainerCleanupStop   chan struct{}
}

// NewServer creates a new instance of the OpenRun Server
func NewServer(config *types.ServerConfig) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setup secrets manager

// Update secrets in the config (including telemetry headers, which are
// resolved before being passed to the OTLP exporter).

// Initialize telemetry after secrets are resolved so OTLP headers can use
// {{ secret ... }} references. A failure here is logged but does not block
// server startup; observability is non-essential.

// Setup OAuth auth

// Setup SAML auth

// no-op, logging is disabled

// if command is empty string, that means either containers are disabled in config or no container command found

// Initialize dynamic config if not already done

// Start the idle shutdown check
// run sync every minute

func (s *Server) GetDynamicConfig() types.DynamicConfig {
	_ = "STUB: not implemented"
	return *new(types.DynamicConfig)
}

// return a copy of the dynamic config

func (s *Server) SaveDynamicConfig(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) updateDynamicConfigCache(ctx context.Context, newConfig *types.DynamicConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// update the rbac config so that it updates its caches

func (s *Server) UpdateDynamicConfig(ctx context.Context, newConfig *types.DynamicConfig, force bool) (*types.DynamicConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// stale update

func (s *Server) appNotifyHandler(updatePayload types.AppUpdatePayload) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) configNotifyHandler(updatePayload types.ConfigUpdatePayload) {
	_ = "STUB: not implemented"
	return
}

// get the latest dynamic config from database

// updateConfigSecrets updates the secrets in the server config using the evalSecret function
func updateConfigSecrets(config *types.ServerConfig, evalSecret func(string) (string, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO : eval store and fs db connections secrets

// handleAppClose listens for app close notifications and removes the app from the store
func (s *Server) handleAppClose() { _ = "STUB: not implemented"; return }

// setupAdminAccount sets up the basic auth password for admin account. If admin user is unset,
// that means admin account is not enabled. If AdminPasswordBcrypt is set, it will be used as
// the password hash for the admin account. If AdminPasswordBcrypt is not set, a random password
// will be generated for that server startup. The generated password will be printed to stdout.
func (s *Server) setupAdminAccount() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Start starts the OpenRun Server
func (s *Server) Start() error { _ = "STUB: not implemented"; return nil }

// Change to OPENRUN_HOME directory, helps avoid length limit on UDS file (around 104 chars)

// Start unix domain socket server

// use relative path

// Unix domain sockets is enabled

// Cannot dial also, so it's safe to delete the socket file

// UDS is admin-only, peer is authenticated by file permissions

//nolint:errcheck

//nolint:errcheck

// Start HTTP and HTTPS servers

// public HTTP listener; do not extract incoming traceparent

// app traffic is traced at the app layer with app redaction policy

// webhook URLs may include secrets in the path

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func (s *Server) setupHTTPSServer() (*http.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Certmagic is enabled

// Use Let's Encrypt staging server

// Use the database backed storage

// Certmagic is disabled, use certs from disk or create self signed ones

// Check if certificate and key files exist on disk for the domain

// If mkcerts is enabled and certificate or key files do not exist, generate them
// Locking is global, not per domain

// If certificate and key files exist, load them

// Request client certificates, verification is done in the handler

// Create a rate-limited error logger for TLS handshake errors

// public HTTPS listener; do not extract incoming traceparent

// app traffic is traced at the app layer with app redaction policy

// webhook URLs may include secrets in the path

func loadRootCAs(rootCertFile string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stop stops the OpenRun Server
func (s *Server) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) GetListAppsApp(ctx context.Context) (*app.App, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) ParseGlob(appGlob string) ([]types.AppInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AuthorizeList checks if the user has access to perform list operation on the specified app
// For RBAC mode, uses RBAC permissions. For non-RBAC mode, look at whether app is using
// same authentication types as used by the caller
func (s *Server) AuthorizeList(userId string, app *types.AppInfo, groups []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// RBAC auth is enabled, verify access

// Admin user is always authorized

// Verify user_id as set in authenticateAndServeApp

// No auth required for this app, authorize access

// Check Oauth provider is the same as the app's provider

// KVInitConstant initializes a constant value in the DB. If the value already exists, it returns the existing value.
// If the value does not exist, it inserts the new value and returns it. If another server inserts the value concurrently,
// it fetches the value from the DB and returns it.
func (s *Server) KVInitConstant(ctx context.Context, keyName string, newValue []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Value already exists in DB, use it

// New value inserted, return it

// Failed to insert, maybe concurrent insert from another server, get the value from the DB

func (s *Server) CleanupVersions() {
	_ = "STUB: not implemented"
	// Cleanup old versions of apps
	return
}

// KVStore is an interface for a key-value store. Implemented by metadata.Metadata
type KVStore interface {
	FetchKV(ctx context.Context, key string) (map[string]any, error)
	FetchKVBlob(ctx context.Context, key string) ([]byte, error)
	StoreKV(ctx context.Context, key string, value map[string]any, expireAt *time.Time) error
	StoreKVBlob(ctx context.Context, key string, value []byte, expireAt *time.Time) error
	UpsertKVBlob(ctx context.Context, key string, value []byte, expireAt *time.Time) error

	UpdateKV(ctx context.Context, key string, value map[string]any) error
	UpdateKVBlob(ctx context.Context, key string, value []byte) error
	DeleteKV(ctx context.Context, key string) error
}
