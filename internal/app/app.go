// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi/v5"
	"github.com/openrundev/openrun/internal/app/action"
	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/app/apptype"
	"github.com/openrundev/openrun/internal/app/dev"
	"github.com/openrundev/openrun/internal/app/starlark_type"
	"github.com/openrundev/openrun/internal/container"
	"github.com/openrundev/openrun/internal/rbac"
	"github.com/openrundev/openrun/internal/types"
	"go.opentelemetry.io/otel/attribute"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"go.starlark.net/syntax"
)

// App is the main object that represents a OpenRun app. It is created when the app is loaded
type App struct {
	*types.Logger
	*types.AppEntry
	Name         string
	CustomLayout bool
	notifyClose  chan<- types.AppPathDomain // Channel to notify server to close the app

	codeConfig       *apptype.CodeConfig
	sourceFS         *appfs.SourceFs
	initMutex        sync.Mutex
	initialized      bool
	reloadError      error
	reloadStartTime  time.Time
	appDev           *dev.AppDev
	appStyle         *dev.AppStyle
	systemConfig     *types.SystemConfig
	storeInfo        *starlark_type.StoreInfo
	paramInfo        map[string]apptype.AppParam
	paramValuesStr   map[string]string   // the param values for the app, from metadata and defaults
	paramDict        starlark.StringDict // the Starlark param values for the app
	plugins          *AppPlugins
	containerHandler *ContainerHandler
	serverConfig     *types.ServerConfig

	globals      starlark.StringDict    // global variables defined in starlark code
	appDef       *starlarkstruct.Struct // app starlark definition
	errorHandler starlark.Callable      // error handler function
	appRouter    *chi.Mux               // router for the app
	actions      []*action.Action       // actions defined for the app

	usesHtmlTemplate bool                          // Whether the app uses HTML templates, false if only JSON APIs
	template         *template.Template            // unstructured templates, no base_templates defined
	templateMap      map[string]*template.Template // structured templates, base_templates defined
	staticOnly       bool                          // app has only static files, no HTML routes
	redirectBarePath bool                          // whether to redirect bare path requests to the full path with trailing slash
	jsLibs           []types.JSLibrary             // JS libraries used by the app

	watcher       *fsnotify.Watcher
	sseListeners  []chan SSEMessage
	funcMap       template.FuncMap
	starlarkCache map[string]*starlarkCacheEntry

	// App config that takes default values from toml config, overridden with app level metadata.
	// It is important that this property is used instead of reading from app metadata config, so that toml
	// config defaults are applied.
	AppConfig types.AppConfig

	lastRequestTime atomic.Int64
	secretEvalFunc  func([][]string, string, string) (string, error)
	auditInsert     func(*types.AuditEvent) error
	AppRunPath      string       // path to the app run directory
	rbacApi         rbac.RBACAPI // the rbac api to use

	// telemetryAttrs caches the immutable per-app OpenTelemetry attributes so
	// that ServeHTTP does not allocate them on every request.
	telemetryAttrs         []attribute.KeyValue
	telemetryIdentityAttrs []attribute.KeyValue

	activeContainerName container.ContainerName
	bindings            []*types.Binding
}

type starlarkCacheEntry struct {
	globals starlark.StringDict
	err     error
}

type SSEMessage struct {
	event string
	data  string
}

func NewApp(sourceFS *appfs.SourceFs, workFS *appfs.WorkFs, logger *types.Logger,
	appEntry *types.AppEntry, systemConfig *types.SystemConfig,
	plugins map[string]types.PluginSettings, appConfig types.AppConfig, notifyClose chan<- types.AppPathDomain,
	secretEvalFunc func([][]string, string, string) (string, error),
	auditInsert func(*types.AuditEvent) error, serverConfig *types.ServerConfig,
	rbacApi rbac.RBACAPI, bindings []*types.Binding) (*App, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *App) Initialize(ctx context.Context, dryRun types.DryRun) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *App) Close() error { _ = "STUB: not implemented"; return nil }

// ActiveContainerName returns the container from the last successful app reload.
func (a *App) ActiveContainerName() (container.ContainerName, bool) {
	_ = "STUB: not implemented"
	return *new(container.ContainerName), false
}

func (a *App) updateActiveContainerNameLocked() { _ = "STUB: not implemented"; return }

func (a *App) ResetFS() { _ = "STUB: not implemented"; return }

func (a *App) Reload(ctx context.Context, force, immediate bool, dryRun types.DryRun, reloadContainer bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Current request is older than the last reloaded request, ignore

// Sleep to allow for multiple file changes to be processed together
// For slower machines, this can be increased, default is 300ms. The tailwind watcher
// especially might need a higher value

// Clear any cached data

// Config lock is not present, use default config

// Config lock file is present, read defaults from that

// Load Starlark config, AppConfig is updated with the settings contents

// Initialize style configuration

// Copy settings into appdev

// Setup the CSS files

// Start the watcher for CSS files unless disabled

// Allow the app to start even if the watcher fails

// Setup the JS libraries

// Create the generated HTML

// Parse HTML templates if there are HTML routes or action uses HTML templates

// No base templates found, use the default unstructured templates

// No html templates found, but app has html routes

// no html templates, ignore error

// Some other error parsing templates, report

// Base templates found, using structured templates

// structured templates are not supported for actions currently

const (
	CONTAINERFILE = "Containerfile"
	DOCKERFILE    = "Dockerfile"
)

func (a *App) loadContainerManager(ctx context.Context, stripAppPath bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Plugin not authorized, skip any container files

// Parse the source file specification

// Look for a file in the source fs (ignoring spec). If not found,
// look in spec files also.

// Containerfile/Dockerfile not found in source, check in spec files also

// Custom container file (or image name prefixed with image:)

func (a *App) executeTemplate(w io.Writer, template, partial string, data any) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *App) loadSchemaInfo(sourceFS *appfs.SourceFs) error {
	_ = "STUB: not implemented"
	// Load the schema info
	return nil
}

// Ignore absence of schema file

func (a *App) loadParamsInfo(sourceFS *appfs.SourceFs) error {
	_ = "STUB: not implemented"
	// Load the params info
	return nil
}

// Ignore absence of params file

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// some apps like gradio need redirect to the full path with trailing slash

// new api call, update last request time

// Pass cached app attrs and per-request attrs as separate WithAttributes
// options. The SDK concatenates them internally, so we avoid the
// per-request allocation of a combined slice and the cached
// a.telemetryAttrs slice is referenced (not copied) here.

func (a *App) startWatcher() error { _ = "STUB: not implemented"; return nil }

// Start listening for events.

// If a reload is in progress, ignore the event

// If a reload has happened recently, ignore the event

//nolint:staticcheck
// ignore chmod events

// Force clients to refresh if reload failed

// Add watcher path.

func (a *App) addSSEClient(newChan chan SSEMessage) { _ = "STUB: not implemented"; return }

func (a *App) removeSSEClient(chanRemove chan SSEMessage) { _ = "STUB: not implemented"; return }

func (a *App) notifyClients() { _ = "STUB: not implemented"; return }

func (a *App) sseHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

//keeping the connection alive with keep-alive protocol

//listen to signal to close and unregister

//nolint:errcheck
//nolint:errcheck

//nolint:errcheck

// loadStarlark loads a starlark file. The main app.star, if it calls load on a file with .star suffix, then
// this function is used to load the starlark file.
func (a *App) loadStarlark(thread *starlark.Thread, module string, cache map[string]*starlarkCacheEntry) (starlark.StringDict, error) {
	_ = "STUB: not implemented"
	return *new(starlark.StringDict), nil
}

// request for package whose loading is in progress

// Add a placeholder to indicate "load in progress".

// Update the cache.

func AppFileOptions() *syntax.FileOptions { _ = "STUB: not implemented"; return nil }

// updateAppConfig updates the app defaults from the metadata
// It creates a TOML intermediate string so that the TOML parsing can be used
func (a *App) updateAppConfig() error { _ = "STUB: not implemented"; return nil }

func (a *App) getSecretsAllowed(plugin, function string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func (a *App) getStarPath(name string) string { _ = "STUB: not implemented"; return "" }
