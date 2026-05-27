// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package dev

import (
	"embed"

	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/app/apptype"
	"github.com/openrundev/openrun/internal/types"
)

//go:embed index_gen.go.html openrun_gen.go.html
var embedHtml embed.FS
var indexEmbed, openrunGenEmbed []byte

func init() {
	var err error
	if indexEmbed, err = embedHtml.ReadFile(apptype.INDEX_GEN_FILE); err != nil {
		panic(err)
	}
	if openrunGenEmbed, err = embedHtml.ReadFile(apptype.CLACE_GEN_FILE); err != nil {
		panic(err)
	}
}

// AppDev is the main object that represents a OpenRun app in dev mode. It is created when the app is loaded with is_dev true
// and handles the styling and js library related functionalities. Access to this is synced through the initMutex in App.
// The reload method in App is the main access point to this object
type AppDev struct {
	*types.Logger

	CustomLayout bool
	Config       *apptype.CodeConfig
	systemConfig *types.SystemConfig
	sourceFS     *appfs.WritableSourceFs
	workFS       *appfs.WorkFs
	AppStyle     *AppStyle

	filesDownloaded map[string][]string
	JsLibs          []types.JSLibrary
	jsCache         map[types.JSLibrary]string
}

func NewAppDev(logger *types.Logger, sourceFS *appfs.WritableSourceFs, workFS *appfs.WorkFs, appStyle *AppStyle, systemConfig *types.SystemConfig) *AppDev {
	_ = "STUB: not implemented"
	return nil
}

// downloadFile downloads the files from the url, unless it was already loaded for this app in the current
// server session.
func (a *AppDev) downloadFile(url string, appFS *appfs.WritableSourceFs, path string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

// SetupJsLibs sets up the js libraries for the app.
func (a *AppDev) SetupJsLibs() error { _ = "STUB: not implemented"; return nil }

// Setup failed and cannot check if file exists, error out

// Setup failed and file does not exist, error out with original error

// Cache that this lib is setup

// This lib is in the cache, but not in current list of libs. Remove it
// from the disk and from cache.

// GenerateHTML generates the default HTML template files for the app.
func (a *AppDev) GenerateHTML() error {
	_ = "STUB: not implemented"
	// The header name of contents have changed, recreate it. Since reload creates the header
	// file and updating the file causes the FS watcher to call reload, we have to make sure the
	// file is updated only if there is an actual content change
	return nil
}

// If generated index file exists, remove it

func (a *AppDev) SaveConfigLockFile() error { _ = "STUB: not implemented"; return nil }

// Close the app dev session
func (a *AppDev) Close() error { _ = "STUB: not implemented"; return nil }
