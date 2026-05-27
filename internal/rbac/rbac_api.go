// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package rbac

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
)

// AuthorizeAny checks if the user has access to any of the specified custom permissions
// Used for app level permissions, like actions access
func (h *RBACManager) AuthorizeAny(ctx context.Context, permissions []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Authorize checks if the user has access to the specified permission
func (h *RBACManager) Authorize(ctx context.Context, permission types.RBACPermission, isCustomPermission bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetCustomPermissions returns the custom permissions for the user on the current app
func (h *RBACManager) GetCustomPermissions(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsAppRBACEnabled checks if the RBAC is enabled for the current app
func (h *RBACManager) IsAppRBACEnabled(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// rbac is not enabled at the config level

// app auth does not have rbac enabled

type RBACAPI interface {
	AuthorizeAny(ctx context.Context, permissions []string) (bool, error)
	Authorize(ctx context.Context, permission types.RBACPermission, isAppLevelPermission bool) (bool, error)
	GetCustomPermissions(ctx context.Context) ([]string, error)
	IsAppRBACEnabled(ctx context.Context) bool
}

var _ RBACAPI = (*RBACManager)(nil)
