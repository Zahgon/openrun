// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package system

import (
	"text/template"
)

// GetFuncMap returns a template.FuncMap that includes all the sprig functions except for env and expandenv.
func GetFuncMap() template.FuncMap { _ = "STUB: not implemented"; return *new(template.FuncMap) }
