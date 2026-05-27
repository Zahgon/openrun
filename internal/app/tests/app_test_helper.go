// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app_test

import (
	"io/fs"
	"strings"
	"text/template"
	"time"

	"github.com/openrundev/openrun/internal/app"
	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/rbac"
	"github.com/openrundev/openrun/internal/types"

	_ "github.com/openrundev/openrun/internal/app/store" // Register db plugin
	_ "github.com/openrundev/openrun/plugins"            // Register builtin plugins
)

func CreateDevModeTestApp(logger *types.Logger, fileData map[string]string) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestApp(logger *types.Logger, fileData map[string]string) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestAppConfig(logger *types.Logger, fileData map[string]string, appConfig types.AppConfig) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestAppParams(logger *types.Logger, fileData map[string]string, params map[string]string) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestAppRoot(logger *types.Logger, fileData map[string]string) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestAppPlugin(logger *types.Logger, fileData map[string]string,
	plugins []string, permissions []types.Permission, pluginConfig map[string]types.PluginSettings) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestAppPluginRoot(logger *types.Logger, fileData map[string]string,
	plugins []string, permissions []types.Permission, pluginConfig map[string]types.PluginSettings) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateDevAppPlugin(logger *types.Logger, fileData map[string]string, plugins []string,
	permissions []types.Permission, pluginConfig map[string]types.PluginSettings) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestAppPluginId(logger *types.Logger, fileData map[string]string,
	plugins []string, permissions []types.Permission, pluginConfig map[string]types.PluginSettings, id string, settings types.AppSettings) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestAppAuthorizer(logger *types.Logger, fileData map[string]string,
	plugins []string, permissions []types.Permission, pluginConfig map[string]types.PluginSettings, rbacApi rbac.RBACAPI) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateTestAppInt(logger *types.Logger, path string, fileData map[string]string, isDev bool,
	plugins []string, permissions []types.Permission, pluginConfig map[string]types.PluginSettings,
	id string, settings types.AppSettings, params map[string]string, appConfig *types.AppConfig,
	rbacApi rbac.RBACAPI) (*app.App, *appfs.WorkFs, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func createTestAppEntry(id, path string, isDev bool, metadata types.AppMetadata) *types.AppEntry {
	_ = "STUB: not implemented"
	return nil
}

type TestReadFS struct {
	fileData map[string]string
}

var _ appfs.ReadableFS = (*TestReadFS)(nil)

type TestWriteFS struct {
	*TestReadFS
}

var _ appfs.WritableFS = (*TestWriteFS)(nil)

type TestFileInfo struct {
	f *TestFile
}

func (fi *TestFileInfo) Name() string { _ = "STUB: not implemented"; return "" }

func (fi *TestFileInfo) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (fi *TestFileInfo) Mode() fs.FileMode { _ = "STUB: not implemented"; return *new(fs.FileMode) }

func (fi *TestFileInfo) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (fi *TestFileInfo) IsDir() bool { _ = "STUB: not implemented"; return false }

func (fi *TestFileInfo) Sys() any { _ = "STUB: not implemented"; return *new(any) }

type TestFile struct {
	name   string
	data   string
	reader *strings.Reader
}

func CreateTestFile(name string, data string) *TestFile { _ = "STUB: not implemented"; return nil }

func (f *TestFile) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (f *TestFile) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *TestFile) Read(dst []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *TestFile) Close() error { _ = "STUB: not implemented"; return nil }

func (f *TestReadFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

func (f *TestReadFS) ReadFile(name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *TestReadFS) Glob(pattern string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *TestReadFS) ParseFS(funcMap template.FuncMap, patterns ...string) (*template.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *TestReadFS) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (f *TestReadFS) StatNoSpec(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func (d *TestReadFS) StaticFiles() []string { _ = "STUB: not implemented"; return nil }

func (d *TestReadFS) FileHash(excludeGlob []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *TestReadFS) CreateTempSourceDir() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *TestReadFS) Reset() {
	_ = "STUB: not implemented"
	// do nothing
	return
}

func (f *TestWriteFS) Write(name string, bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (f *TestWriteFS) Remove(name string) error { _ = "STUB: not implemented"; return nil }
