// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/openrundev/openrun/internal/types"
	"github.com/urfave/cli/v2"
)

func getServerCommands(serverConfig *types.ServerConfig, clientConfig *types.ClientConfig) ([]*cli.Command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func startServer(cCtx *cli.Context, serverConfig *types.ServerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// no profiling

// block forever, profiling will exit on interrupt

// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)

// Block until we receive our signal.

// Create a deadline to wait for.

func stopServer(_ *cli.Context, clientConfig *types.ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func showConfig(_ *cli.Context, clientConfig *types.ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func updateConfig(cCtx *cli.Context, clientConfig *types.ClientConfig) error {
	_ = "STUB: not implemented"
	return nil
}
