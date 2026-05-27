// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package system

import (
	"embed"

	"github.com/openrundev/openrun/internal/types"
)

const DEFAULT_CONFIG = "openrun.default.toml"

//go:embed "openrun.default.toml"
var f embed.FS

func getEmbeddedToml() (string, error) { _ = "STUB: not implemented"; return "", nil }

//nolint:errcheck

// NewServerConfigEmbedded reads the embedded toml file and creates a ServerConfig
func NewServerConfigEmbedded() (*types.ServerConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadServerConfig loads a ServerConfig from the given contents
func LoadServerConfig(contents string, config *types.ServerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// NewClientConfigEmbedded reads the embedded toml file and creates a ClientConfig
func NewClientConfigEmbedded() (*types.ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadClientConfig load a ClientConfig from the given contents
func LoadClientConfig(contents string, config *types.ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadGlobalConfig load a GlobalConfig from the given contents
func LoadGlobalConfig(contents string, config *types.GlobalConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func GetDefaultConfigs() (*types.GlobalConfig, *types.ClientConfig, *types.ServerConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
