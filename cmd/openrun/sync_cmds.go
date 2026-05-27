// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/openrundev/openrun/internal/types"
	"github.com/urfave/cli/v2"
)

func initSyncCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func syncScheduleCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func syncListCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func syncRunCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func syncDeleteCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func printSyncList(cCtx *cli.Context, sync []*types.SyncEntry, format string) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func getSyncType(sync *types.SyncEntry) string { _ = "STUB: not implemented"; return "" }
