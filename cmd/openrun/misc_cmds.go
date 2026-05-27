// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/openrundev/openrun/internal/types"
	"github.com/urfave/cli/v2"
)

func getPasswordCommands(clientConfig *types.ClientConfig) ([]*cli.Command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generatePassword(cCtx *cli.Context) error { _ = "STUB: not implemented"; return nil }

func promptPassword(prompt string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func readPassword() (string, error) { _ = "STUB: not implemented"; return "", nil }

func printVersion(cCtx *cli.Context) { _ = "STUB: not implemented"; return }
