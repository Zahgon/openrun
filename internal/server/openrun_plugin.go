// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"go.starlark.net/starlark"
)

func initOpenRunPlugin(server *Server) { _ = "STUB: not implemented"; return }

type openrunPlugin struct {
	server *Server
}

func (c *openrunPlugin) ListAllApps(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (c *openrunPlugin) ListApps(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (c *openrunPlugin) listAppsImpl(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple, permCheck bool, apiName string) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

//nolint:errcheck

// Filter out internal apps

// Check query filter

// If path glob is specified, check if app matches. If internal apps are to be included,
// check if main app matches

//nolint:errcheck

// Last path, no glob
//nolint:errcheck

//nolint:errcheck

func getSourceUrl(sourceUrl, branch string) string { _ = "STUB: not implemented"; return "" }

func (c *openrunPlugin) ListAuditEvents(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// Postgres

// Postgres

//nolint:errcheck

//nolint:errcheck
func (c *openrunPlugin) ListOperations(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func getOpList(op string) ([]any, string) { _ = "STUB: not implemented"; return nil, "" }

// Some infrequent operations like account link are not included in the list for now

func (c *openrunPlugin) ListSync(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

//nolint:errcheck
