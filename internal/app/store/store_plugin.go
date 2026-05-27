// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"database/sql"

	"github.com/openrundev/openrun/internal/app"
	"github.com/openrundev/openrun/internal/plugin"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

const (
	TRANSACTION_KEY = "transaction"
)

func init() {
	h := &storePlugin{}
	pluginFuncs := []plugin.PluginFunc{
		app.CreatePluginApi(h.Begin, app.READ),
		app.CreatePluginApi(h.Commit, app.WRITE),
		app.CreatePluginApi(h.Rollback, app.READ),

		app.CreatePluginApiName(h.SelectById, app.READ, "select_by_id"),
		app.CreatePluginApi(h.Select, app.READ),
		app.CreatePluginApiName(h.SelectOne, app.READ, "select_one"),
		app.CreatePluginApi(h.Count, app.READ),
		app.CreatePluginApi(h.Insert, app.WRITE),
		app.CreatePluginApi(h.Update, app.WRITE),
		app.CreatePluginApiName(h.DeleteById, app.WRITE, "delete_by_id"),
		app.CreatePluginApi(h.Delete, app.WRITE),
	}
	app.RegisterPlugin("store", NewStorePlugin, pluginFuncs)
}

type storePlugin struct {
	sqlStore *SqlStore
}

func NewStorePlugin(pluginContext *types.PluginContext) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func fetchTransation(thread *starlark.Thread) *sql.Tx { _ = "STUB: not implemented"; return nil }

func (s *storePlugin) Begin(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) Commit(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) Rollback(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) Insert(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) SelectById(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) Update(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) DeleteById(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) SelectOne(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

type filterData struct {
	data map[string]any
}

func (e *filterData) Unpack(value starlark.Value) error { _ = "STUB: not implemented"; return nil }

func (s *storePlugin) Select(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) Count(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (s *storePlugin) Delete(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}
