// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

func (a *App) Audit() (*types.ApproveResult, error) { _ = "STUB: not implemented"; return nil, nil }

// Load the starlark file rather than the plugin

// The loader in audit mode is used to track the modules that are loaded.
// A copy of the real loader's response is returned, with builtins replaced with dummy methods,
// so that the audit can be run without any side effects

// Replace all the builtins with dummy methods

// TODO use logger

// This runs the starlark script, with dummy plugin methods
// The intent is to load the permissions from the app definition while trying
// to avoid any potential side effects from script

func needsApproval(a *types.ApproveResult) bool { _ = "STUB: not implemented"; return false }

//TODO: sort slices before checking equality

// needsApprovalWithServerConfig re-evaluates whether approval is needed after accounting
// for loads and permissions that are already covered by the server config.
func needsApprovalWithServerConfig(a *types.ApproveResult, serverPerms []types.Permission) bool {
	_ = "STUB: not implemented"
	return false
}

// Filter new loads: remove any that are covered by server config or already approved

// Filter new permissions: remove any that are covered by server config or already approved

// Check if it's in the approved list

// permissionCoveredByServerConfig checks if a permission declared by an app is already
// covered by a server config permission entry. Server config permissions can use regex
// for arguments, so this does regex-aware matching.
func permissionCoveredByServerConfig(perm types.Permission, serverPerms []types.Permission) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *App) createApproveResponse(loads []string, globals starlark.StringDict) (*types.ApproveResult, error) {
	_ = "STUB: not implemented"
	// the App entry should not get updated during the audit call, since there
	// can be audit calls when the app is running.
	return nil, nil
}

// permission order needs to match for now
