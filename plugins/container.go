// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package plugins

import (
	"github.com/openrundev/openrun/internal/app"
	"github.com/openrundev/openrun/internal/app/apptype"
	"github.com/openrundev/openrun/internal/plugin"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

func init() {
	h := &containerPlugin{}
	pluginFuncs := []plugin.PluginFunc{
		app.CreatePluginApi(h.Config, app.READ), // config API
		app.CreatePluginApi(h.Run, app.READ_WRITE),
		app.CreatePluginConstant("URL", starlark.String(apptype.CONTAINER_URL)),
		app.CreatePluginConstant("AUTO", starlark.String(types.CONTAINER_SOURCE_AUTO)),
		app.CreatePluginConstant("NIXPACKS", starlark.String(types.CONTAINER_SOURCE_NIXPACKS)),
		app.CreatePluginConstant("IMAGE_PREFIX", starlark.String(types.CONTAINER_SOURCE_IMAGE_PREFIX)),
		app.CreatePluginConstant("COMMAND", starlark.String(types.CONTAINER_LIFETIME_COMMAND)),
	}
	app.RegisterPlugin("container", NewContainerPlugin, pluginFuncs)
}

type containerPlugin struct {
	pluginContext *types.PluginContext
}

func NewContainerPlugin(pluginContext *types.PluginContext) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *containerPlugin) Run(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (c *containerPlugin) Config(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}
