// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package dev

import (
	"regexp"

	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/types"
)

var npmPackageNamePattern = regexp.MustCompile(`^(?:@[A-Za-z0-9][A-Za-z0-9._~-]*/)?[A-Za-z0-9][A-Za-z0-9._~-]*(?:/[A-Za-z0-9][A-Za-z0-9._~-]*)*$`)

func NewLibrary(url string) *types.JSLibrary { _ = "STUB: not implemented"; return nil }

func NewLibraryESM(packageName string, version string, esbuildArgs []string) *types.JSLibrary {
	_ = "STUB: not implemented"
	return nil
}

type JsLibManager struct {
	types.JSLibrary
}

func (j *JsLibManager) Setup(dev *AppDev, sourceFS *appfs.WritableSourceFs, workFS *appfs.WorkFs) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
	//nolint:staticcheck
}

func (j *JsLibManager) setupEsbuild(dev *AppDev, sourceFS *appfs.WritableSourceFs, workFS *appfs.WorkFs) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Parse the build options from the esbuild args

//options.AbsWorkingDir = sourceFS.Root this fails if the source dir is not absolute

// Add node paths to the esbuild options to customize the node_module location

// Run esbuild to generate the output file

// Return the target file name. The caller can check if the file exists to determine if the
// setup was successful even though this step failed

func (j *JsLibManager) generateSourceFile(workFS *appfs.WorkFs) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func sanitizeFileName(input string) string { _ = "STUB: not implemented"; return "" }

func validatePackageName(name string) error { _ = "STUB: not implemented"; return nil }
