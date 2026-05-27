// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package rbac

import (
	"regexp"
	"sync"

	"github.com/openrundev/openrun/internal/types"
)

const RBAC_AUTH_PREFIX = "rbac:"
const RBAC_GROUP_PREFIX = "group:"
const RBAC_ROLE_PREFIX = "role:"
const RBAC_CUSTOM_PREFIX = "custom:" // used for app level custom permissions
const RBAC_REGEX_PREFIX = "regex:"   // used for regex matching in users list

type RBACManager struct {
	*types.Logger
	RbacConfig   *types.RBACConfig
	serverConfig *types.ServerConfig
	mu           sync.RWMutex

	groups      map[string][]string               // group name to user ids (with group hierarchy resolved)
	roles       map[string][]types.RBACPermission // role name to permissions (with role: hierarchy resolved)
	regexCache  map[string]*regexp.Regexp         // cache of compiled regex patterns
	customPerms []string                          // custom permissions are permissions defined by the user. This list does not have the custom: prefix
}

func NewRBACHandler(logger *types.Logger, rbacConfig *types.RBACConfig, serverConfig *types.ServerConfig) (*RBACManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *RBACManager) AuthorizeInt(user string, appPathDomain types.AppPathDomain,
	appAuthSetting string, permission types.RBACPermission, groups []string, isAppLevelPermission bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// rbac is not enabled, authorize all requests

// admin user is always authorized if enabled

// if app auth does not have rbac enabled, authorize access for Access permission
// If authenticated, then app access is allowed

// Trim stage and preview suffixes, grant check are done on the main app path

// GetCustomPermissions returns the custom permissions set for the user for the given app path domain and app auth setting
// Values in returned list do not have the custom: prefix
func (h *RBACManager) GetCustomPermissionsInt(user string, appPathDomain types.AppPathDomain, appAuthSetting string,
	groups []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// rbac is not enabled, authorize all requests

// admin user is always authorized if enabled

func (h *RBACManager) checkGrants(inputUser string, appPathDomain types.AppPathDomain,
	inputPermission types.RBACPermission, groups []string, isAppLevelPermission bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// User, role and target matched. This is a valid grant.

func (h *RBACManager) checkGrant(grant types.RBACGrant, inputUser string, appPathDomain types.AppPathDomain,
	inputPermission types.RBACPermission, groups []string, isAppLevelPermission bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// granted group name matched group as found from SSO login

// Check for direct user match

// Check for regex patterns in the group

// user in grant  is a regex, match it against the input user

// user matched, check if role matches

// app level permission, look for grant with custom: prefix

func (h *RBACManager) initGroupInfo(rbacConfig *types.RBACConfig) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize all groups

// Helper function to recursively resolve group membership

// Resolve all groups

func (h *RBACManager) initRoleInfo(rbacConfig *types.RBACConfig) (map[string][]types.RBACPermission, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize all roles

// Helper function to recursively resolve role permissions

// Resolve all roles

// Keep track of all custom perms (deduplicated)

// Convert map to slice

func (h *RBACManager) validateGrants(rbacConfig *types.RBACConfig) error {
	_ = "STUB: not implemented"
	// Skip validation if RBAC is disabled
	return nil
}

// groups can be passed dynamically (for SSO login), so we don't need to validate them
// Validate role references in Roles

func (h *RBACManager) UpdateRBACConfig(rbacConfig *types.RBACConfig) error {
	_ = "STUB: not implemented"
	return nil
}
