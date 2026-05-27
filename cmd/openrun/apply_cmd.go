// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/openrundev/openrun/internal/types"
	"github.com/urfave/cli/v2"
)

func initApplyCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func printApplyResponse(cCtx *cli.Context, applyResponse *types.AppApplyResponse) {
	_ = "STUB: not implemented"
	return
}

// Server does not return these for reload to reduce the noise
