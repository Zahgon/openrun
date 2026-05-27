// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/openrundev/openrun/internal/types"
	"github.com/urfave/cli/v2"
)

func initWebhookCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func webhookListCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func printWebhookList(cCtx *cli.Context, tokens []types.AppToken, format string) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func webhookCreateCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func webhookDeleteCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}
