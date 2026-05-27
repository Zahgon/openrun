// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"sync"

	"github.com/openrundev/openrun/internal/types"
)

// AdminBasicAuth implements basic auth for the admin user account.
// Cache the success auth header to avoid the bcrypt hash check penalty
// Basic auth is supported for admin user only, and changing it requires service restart.
// Caching the sha of the successful auth header allows us to skip the bcrypt check
// which significantly improves performance.
type AdminBasicAuth struct {
	*types.Logger
	config *types.ServerConfig

	mu            sync.RWMutex
	authShaCached string
}

func NewAdminBasicAuth(logger *types.Logger, config *types.ServerConfig) *AdminBasicAuth {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdminBasicAuth) authenticate(authHeader string) bool {
	_ = "STUB: not implemented"
	return false
}

// slow down brute force attacks

// Cached header matches, so we can skip the rest of the auth checks

// slow down brute force attacks

// Successful request, so we can cache the auth header

func (a *AdminBasicAuth) BasicAuth(authHeader string) (username, password string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func (a *AdminBasicAuth) parseBasicAuth(auth string) (username, password string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}
