// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/openrundev/openrun/internal/types"
	"github.com/urfave/cli/v2"
)

const (
	DRY_RUN_FLAG    = "dry-run"
	DRY_RUN_ARG     = "dryRun"
	DRY_RUN_MESSAGE = "\n" + YELLOW + "*** dry-run mode, changes have NOT been committed. ***" + RESET + "\n"
	PATH_SPEC_HELP  = `The (optional) domain and path are separated by a ":". appPathGlob supports a glob pattern.
In the glob, * matches any number of characters, ** matches any number of characters including /.
all is a shortcut for "*:**", which matches all apps across all domains, including no domain.
To prevent shell expansion for *, placing the path in quotes is recommended.
"all" matches all apps across all domains, same as "*:**".
`
	PROMOTE_FLAG = "promote"
	PROMOTE_ARG  = "promote"
)

func initAppCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func dryRunFlag() *cli.BoolFlag { _ = "STUB: not implemented"; return nil }

func appCreateCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

// value can be empty string

func printCreateResult(cCtx *cli.Context, createResult types.AppCreateResponse) {
	_ = "STUB: not implemented"
	return
}

func appListCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func printAppList(cCtx *cli.Context, apps []types.AppResponse, format string) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func appType(app types.AppResponse) string { _ = "STUB: not implemented"; return "" }

func authType(app types.AppResponse) string { _ = "STUB: not implemented"; return "" }

func permType(perm types.Permission) string { _ = "STUB: not implemented"; return "" }

func printApproveResult(approveResult types.ApproveResult) { _ = "STUB: not implemented"; return }

func appDeleteCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func appApproveCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func appReloadCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

// Server does not return these for reload to reduce the noise

func appPromoteCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}
