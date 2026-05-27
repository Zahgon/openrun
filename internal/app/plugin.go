// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"sync"

	"github.com/openrundev/openrun/internal/app/apptype"
	"github.com/openrundev/openrun/internal/plugin"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

type PluginFunctionType int

const (
	READ PluginFunctionType = iota
	WRITE
	READ_WRITE
)

var (
	loaderInitMutex sync.Mutex
	builtInPlugins  map[string]plugin.PluginMap
)

func init() {
	builtInPlugins = make(map[string]plugin.PluginMap)
	initFS()
}

// RegisterPlugin registers a plugin with OpenRun
func RegisterPlugin(name string, builder plugin.NewPluginFunc, funcs []plugin.PluginFunc) {
	_ = "STUB: not implemented"
	return
}

type StarlarkFunction func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

// pluginErrorWrapper wraps the plugin function call with error handling code. If the plugin function returns an error,
// it is wrapped in a PluginResponse. If the starlark function returns a PluginResponse, it is returned as is. Returning
// a error causes the starlark interpreter to panic, so this wrapper is needed to handle the error and return a value which
// the starlark code can handle
func pluginErrorWrapper(f StarlarkFunction, errorHandler starlark.Callable) StarlarkFunction {
	_ = "STUB: not implemented"
	return *new(StarlarkFunction)
}

// Wrap the plugin function call with error handling

// If the return value is already of type PluginResponse, return it without wrapping it

// Update the thread local error state

// Error response wrapped in a PluginResponse

// Success response, wrapped in a PluginResponse

func CreatePluginConstant(name string, value starlark.Value) plugin.PluginFunc {
	_ = "STUB: not implemented"
	return *new(plugin.PluginFunc)
}

func CreatePluginApi(f StarlarkFunction, opType PluginFunctionType) plugin.PluginFunc {
	_ = "STUB: not implemented"
	return *new(plugin.PluginFunc)
}

// -fm denotes function value

// CreatePluginApiName creates a OpenRun plugin function
func CreatePluginApiName(
	f func(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error),
	opType PluginFunctionType,
	name string) plugin.PluginFunc {
	_ = "STUB: not implemented"
	return *new(plugin.PluginFunc)
}

// -fm denotes function value

func GetContext(thread *starlark.Thread) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SavePluginState saves a value in the thread local for the plugin
func SavePluginState(thread *starlark.Thread, key string, value any) {
	_ = "STUB: not implemented"
	return
}

// FetchPluginState fetches a value from the thread local for the plugin
func FetchPluginState(thread *starlark.Thread, key string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// DeferCleanup defers a close function to call when the API handler is done
func DeferCleanup(thread *starlark.Thread, key string, deferFunc apptype.DeferFunc, strict bool) {
	_ = "STUB: not implemented"
	return
}

// ClearCleanup clears a defer function from the thread local
func ClearCleanup(thread *starlark.Thread, key string) { _ = "STUB: not implemented"; return }

// loader is the starlark loader function
func (a *App) loader(thread *starlark.Thread, moduleFullPath string) (starlark.StringDict, error) {
	_ = "STUB: not implemented"
	return *new(starlark.StringDict), nil
}

// Load the starlark file rather than the plugin

// Add calls to the hook function, which will do the permission checks at invocation time to
// verify if the application has approval to call the specified function.
// The audit loader will replace the builtins with dummy methods, so the hook is not added for the audit loader

func parseModulePath(moduleFullPath string) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// pluginLookup looks up the plugin. Audit checks need to be done by the caller
func (a *App) pluginLookup(_ *starlark.Thread, module string) (plugin.PluginMap, error) {
	_ = "STUB: not implemented"
	return *new(plugin.PluginMap), nil
}

// TODO extend loading

func (a *App) pluginHook(appPath, modulePath, accountName, functionName string, pluginInfo *plugin.PluginInfo) *starlark.Builtin {
	_ = "STUB: not implemented"
	return nil
}

// Add the server config allowed permissions to the list

// App has full access to all plugins

// All secrets are allowed

// Get the plugin from the app config

// Get the plugin function using reflection

// Wrap the plugin function call with error handling

// Pass the module full path as a thread local

// Evaluate secrets passed as arguments

//nolint:errcheck

//nolint:errcheck

// Evaluate secrets passed as keyword arguments

//nolint:errcheck

//nolint:errcheck

// Call the builtin function

// Plugin spans are gated separately because data-heavy apps may issue
// many plugin calls per request; we also require an existing parent
// context (set by the surrounding handler) so a missing TL_CONTEXT
// surfaces as no span rather than a silently-detached root span.

func checkPermissions(a *App, modulePath string, functionName string, args starlark.Tuple, pluginInfo *plugin.PluginInfo, permsList []types.Permission) (error, [][]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, nil, false, nil
}

// More arguments than approved are permitted. Also, using kwargs is not allowed for args which are approved

// This permission is not approved, but there may be others which are

// Permission defines isRead, use that

// Use the plugin defined isRead value

// Write API, check if stage/preview has write access

// the secrets this plugin call is allowed access to
