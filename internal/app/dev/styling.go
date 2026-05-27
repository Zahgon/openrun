// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package dev

import (
	"os"
	"os/exec"

	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlarkstruct"
)

const (
	STYLE_FILE_PATH = "static/gen/css/style.css"
)

const (
	TailwindCSS types.StyleType = "tailwindcss"
	DaisyUI     types.StyleType = "daisyui"
	Other       types.StyleType = "other"
	None        types.StyleType = ""
)

// AppStyle is the style related configuration and state for an app. It is created
// when the App is loaded. It keeps track of the watcher process required to rebuild the
// CSS file when the tailwind/daisy config changes. The reload mutex lock in App is used to
// ensure only one call to the watcher is done at a time, no locking is implemented in AppStyle
type AppStyle struct {
	appId          types.AppId
	library        types.StyleType
	themes         []string
	libraryUrl     string
	DisableWatcher bool
	watcher        *exec.Cmd
	watcherState   *WatcherState
	watcherStdout  *os.File
	Light          string
	Dark           string
}

// WatcherState is the state of the watcher process as of when it was last started.
type WatcherState struct {
	library           types.StyleType
	templateLocations []string
}

// Init initializes the AppStyle object from the app definition
func (s *AppStyle) Init(appId types.AppId, appDef *starlarkstruct.Struct) error {
	_ = "STUB: not implemented"
	return nil
}

// No style defined

// Setup sets up the style library for the app. This is called when the app is reloaded.
func (s *AppStyle) Setup(dev *AppDev) error { _ = "STUB: not implemented"; return nil }

// Empty out the style.css file

// Generate the tailwind/daisyui config files

// Download style.css from url

const (
	// TODO: allow custom config file to be specified
	TAILWIND_CONFIG_FILE     = "tailwind.config.js"
	TAILWIND_CONFIG_CONTENTS = `
	module.exports = {
		content: [%s],
		theme: {
		  extend: {},
		},
	  
		plugins: [
		  %s
		],
		%s
	}`

	TAILWIND_INPUT_CONTENTS = `
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
	`
)

func (s *AppStyle) setupTailwindConfig(templateLocations []string, sourceFS *appfs.WritableSourceFs, workFS *appfs.WorkFs) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the action templates to the input list

// File already exists, skip

// Add the template locations to the input list

// StartWatcher starts the watcher process for the app. This is called when the app is reloaded.
func (s *AppStyle) StartWatcher(dev *AppDev) error { _ = "STUB: not implemented"; return nil }

// If config is being switched from tailwind/daisy to other/none, stop any current watcher

func (s *AppStyle) startTailwindWatcher(templateLocations []string, sourceFS *appfs.WritableSourceFs, workFS *appfs.WorkFs, systemConfig *types.SystemConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: log

// TODO: log

// TODO: log

// Since the watcher process creates the file, the unit test framework (in memory filesystem)
// can't be used to test the watcher functionality)

// TODO: log

// Setup stdin/stdout for watcher process

// Start watcher process, wait async for it to complete

// // ensure process group

// this seems to be required for the process to start

// TODO: log

func (s *AppStyle) StopWatcher() error { _ = "STUB: not implemented"; return nil }

func (s *AppStyle) GetStyleType() types.StyleType {
	_ = "STUB: not implemented"
	return *new(types.StyleType)
}
