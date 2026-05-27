// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package rbac

import (
	"github.com/openrundev/openrun/internal/types"
)

// createPathDomain creates a slice of AppPathDomain from a slice of AppInfo
func createPathDomain(apps []types.AppInfo) []types.AppPathDomain {
	_ = "STUB: not implemented"
	return nil
}

// ParseGlobFromInfo parses a path spec in the format of domain:path.  If domain is not specified, it will match empty domain.
// glob patters are supported, *:** matches all apps.
func ParseGlobFromInfo(appPathGlob string, apps []types.AppInfo) ([]types.AppInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MatchGlob(appPathGlob string, app types.AppPathDomain) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ParseGlob parses a path spec in the format of domain:path. If domain is not specified, it will match empty domain.
// glob patters are supported, *:** matches all apps.
func ParseGlob(appPathGlob string, apps []types.AppPathDomain) ([]types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// all apps match

//nolint:staticcheck
