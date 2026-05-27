// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"sync"

	"github.com/openrundev/openrun/internal/plugin"
	"github.com/openrundev/openrun/internal/types"
)

type AppPlugins struct {
	sync.Mutex
	plugins map[string]any

	app          *App
	pluginConfig map[string]types.PluginSettings // pluginName -> accountName -> PluginSettings, from openrun.toml
	accountMap   map[string]string               // pluginName -> accountName, from app account links
}

func NewAppPlugins(app *App, pluginConfig map[string]types.PluginSettings, appAccounts []types.AccountLink) *AppPlugins {
	_ = "STUB: not implemented"
	return nil
}

func (p *AppPlugins) GetPlugin(pluginInfo *plugin.PluginInfo, accountName string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Already initialized, use that

// If account name is specified, use that to lookup the account map

// store.in#myaccount

// If it is just account name, make it full plugin path
