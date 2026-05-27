// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package plugins

import (
	"github.com/openrundev/openrun/internal/app"
	"github.com/openrundev/openrun/internal/plugin"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

func init() {
	h := &proxyPlugin{}
	pluginFuncs := []plugin.PluginFunc{
		app.CreatePluginApi(h.Config, app.READ), // config API, preview/stage permission checks happen in the reverse proxy wrapper
	}
	app.RegisterPlugin("proxy", NewProxyPlugin, pluginFuncs)
}

type proxyPlugin struct {
}

func NewProxyPlugin(pluginContext *types.PluginContext) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (h *proxyPlugin) Config(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}
