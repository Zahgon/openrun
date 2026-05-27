// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package system

import (
	"context"
)

func LoadProperties(filename string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// propertiesSecretProvider is a secret provider that reads secrets from a properties file ( a = b format)
type propertiesSecretProvider struct {
	props map[string]string
}

func (e *propertiesSecretProvider) Configure(ctx context.Context, conf map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *propertiesSecretProvider) GetSecret(ctx context.Context, secretName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *propertiesSecretProvider) GetJoinDelimiter() string { _ = "STUB: not implemented"; return "" }

var _ secretProvider = &propertiesSecretProvider{}

func FileExists(filename string) bool { _ = "STUB: not implemented"; return false }

// any error is treated as not exists

// FindExec looks for the executable in the system PATH and returns its full path.
// If not found, it checks common Homebrew locations for the executable.
func FindExec(name string) string { _ = "STUB: not implemented"; return "" }

// Homebrew services do not have the path set, lookup common paths
