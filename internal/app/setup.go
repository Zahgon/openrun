// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

func (a *App) loadStarlarkConfig(ctx context.Context, dryRun types.DryRun, reloadContainer bool) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO use logger

// Update the app config with entries loaded from the settings map

// Load container config. The proxy config in routes depends on this being loaded first

// Container handler is present, reload the container

// In prod mode, reload when initializing an app. Image-spec apps
// also need to reload on every admin reload so that the upstream
// image digest is resolved and the container is recreated if the
// tag has moved; build-spec apps rely on the source-content hash
// to detect changes and so only need reload on Initialize.

// Initialize the router configuration

func (a *App) createBuiltin() (starlark.StringDict, error) {
	_ = "STUB: not implemented"
	return *new(starlark.StringDict), nil
}

func (a *App) addSchemaTypes(builtin starlark.StringDict) (starlark.StringDict, error) {
	_ = "STUB: not implemented"
	return *new(starlark.StringDict), nil
}

// Create a copy of the builtins, don't modify the original

// Add type module for referencing type names

// Add table module for referencing table names

func (a *App) addParams(builtin starlark.StringDict) (starlark.StringDict, error) {
	_ = "STUB: not implemented"
	return *new(starlark.StringDict), nil
}

// Create a copy of the builtins, don't modify the original

// Add param module for referencing param values

// Set the default value in the paramMap (in the string format)

// no custom value specified

// add additional param values to paramMap

func verifyConfig(globals starlark.StringDict) (*starlarkstruct.Struct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkAppPathStripping checks if the app path should be stripped from the request path for container proxying
// This is required for container health checks.
func (a *App) checkAppPathStripping() (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Iterate through all the routes
		nil
}

// "config" is defined, this must be a proxy config instead of a page definition

// Not proxying to container url, ignore

func fileExists(fs appfs.ReadableFS, name string) bool { _ = "STUB: not implemented"; return false }

func (a *App) initRouter() error { _ = "STUB: not implemented"; return nil }

// Iterate through all the routes

// Root wildcard path, static files are not served

// Single file app

// All app files are served at the root level

// Mount static dir

func (a *App) initActions(router *chi.Mux) error { _ = "STUB: not implemented"; return nil }

// Set the links for all actions

func (a *App) addAction(count int, val starlark.Value, router *chi.Mux) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// addStaticRoot adds the static root directory contents to the router
// Files can be referenced by /<filename>, without /static or /static_root
func (a *App) addStaticRoot(router *chi.Mux) error { _ = "STUB: not implemented"; return nil }

func (a *App) addRoute(count int, router *chi.Mux, routeVal starlark.Value, defaultHandler starlark.Callable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// "config" is defined, this must be a proxy config instead of a page definition

// "full" is not defined, this must be a API route instead of a html route

// Use app level default handler, which could also be nil

// getProxyConfig extracts the proxy config from the proxy definition
func getProxyConfig(count int, proxyDef *starlarkstruct.Struct) (starlark.HasAttrs, error) {
	_ = "STUB: not implemented"
	return *new(starlark.HasAttrs), nil
}

/*
	if config.String() != "ProxyConfig" {
		return nil, fmt.Errorf("proxy entry %d:%s is not a proxy config", count, pathStr)
	}*/

func (a *App) addProxyConfig(count int, router *chi.Mux, proxyDef *starlarkstruct.Struct) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Root wildcard path, static files are not served

// proxying to container url

// To support WebSockets, we need to ensure that the `Connection`, `Upgrade`
// and `Host` headers are forwarded as-is and not modified.

// Set the Host header to target url for non-WebSocket requests, unless
// disabled in proxy config

// If write API, check if preview/stage app is allowed access

// Add X-Openrun- headers to request
// Add the user and custom permissions to the request headers

// Set the response headers

// use the reverse proxy to handle the request

func (a *App) addAPIRoute(basePath string, router *chi.Mux, apiDef *starlarkstruct.Struct, defaultHandler starlark.Callable) error {
	_ = "STUB: not implemented"
	return nil
}

// Use app level default handler, which could also be nil

func (a *App) handleFragments(router *chi.Mux, pagePath string, pageCount int, htmlFile string, block string, page *starlarkstruct.Struct, handlerCallable starlark.Callable) error {
	_ = "STUB: not implemented"
	// Iterate through all the pages
	return nil
}

// No fragments defined

// "partial" is not defined, this must be a API route instead of a html route

// Inherit page level block setting

// If new handler is defined at fragment level, that is verified. Otherwise the
// page level handler is used

func (a *App) createInternalRoutes(router *chi.Mux) error { _ = "STUB: not implemented"; return nil }

func (a *App) userFileHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// this cannot be overridden at the app level

// Check if the user is authorized to access the file

// For app level visibility, the file is accessible if the API is accessible

//nolint:errcheck

// Copy the file content to the response writer
// This streams the content and uses chunked transfer encoding

func (a *App) loadLibraryInfo() ([]types.JSLibrary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No libraries defined
