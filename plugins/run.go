// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package plugins

import (
	"github.com/openrundev/openrun/internal/app"
	"go.starlark.net/starlark"
)

func execCommand(containerHandler *app.ContainerHandler, thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// cwd is not supported in container mode

//nolint:errcheck

// Stream the output to the client using RangeFunc

//nolint:errcheck

//nolint:errcheck

// if no lines in stdout and there was an error (processPartial case), return the error
