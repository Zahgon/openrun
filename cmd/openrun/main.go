// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/openrundev/openrun/internal/system"
	"github.com/openrundev/openrun/internal/types"
)

const configFileFlagName = "config-file"

func getAllCommands(clientConfig *types.ClientConfig, serverConfig *types.ServerConfig) ([]*cli.Command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func globalFlags(globalConfig *types.GlobalConfig) ([]cli.Flag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getConfigPath returns the path to the config file and the home directory
// Uses OPENRUN_HOME env if set. Otherwise uses binaries parent path. Setting OPENRUN_HOME is
// the easiest way to configure. Uses some extra heuristics to help avoid having to setup
// OPENRUN_HOME in the env, by using the binaries parent folder as the default.
// On mac, looks for brew install locations also.
func getConfigPath(cCtx *cli.Context) (string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

// Found OPENRUN_HOME

// OPENRUN_HOME not set and config file is set, use config dir path as OPENRUN_HOME

// Found bin directory, use its parent

// Config file found in parent directory of the executable, use that as path
// To avoid clobbering /usr, check if the path contains the string openrun/clhome

// Running `brew --prefix` would be another option
//nolint:staticcheck
// brew OSX specific checks

// brew linux specific checks

// Linux system level installation

func parseConfig(cCtx *cli.Context, globalConfig *types.GlobalConfig, clientConfig *types.ClientConfig, serverConfig *types.ServerConfig) error {
	_ = "STUB: not implemented"
	// Find OPENRUN_HOME and config file, update OPENRUN_HOME in env
	return nil
}

//nolint:errcheck

//fmt.Fprintf(os.Stderr, "Loading config file: %s, clHome %s\n", filePath, clHome)

func main() {
	globalConfig, clientConfig, serverConfig, err := system.GetDefaultConfigs()
	if err != nil {
		log.Fatal(err)
	}
	globalFlags, err := globalFlags(globalConfig)
	if err != nil {
		log.Fatal(err)
	}
	allCommands, err := getAllCommands(clientConfig, serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:                 "openrun",
		Usage:                "OpenRun client and server https://openrun.dev/",
		EnableBashCompletion: true,
		Suggest:              true,
		Flags:                globalFlags,
		Before: func(ctx *cli.Context) error {
			err := parseConfig(ctx, globalConfig, clientConfig, serverConfig)
			if ctx.Command != nil && ctx.Args().Len() > 0 && ctx.Args().Get(0) == "password" {
				// For password command, ignore error parsing config
				return nil
			}
			if err != nil {
				return fmt.Errorf("error parsing config: %w", err)
			}
			return nil
		},
		ExitErrHandler: func(c *cli.Context, err error) {
			if err != nil {
				fmt.Fprintf(cli.ErrWriter, RED+"error: %s\n"+RESET, err) //nolint:errcheck
				os.Exit(1)
			}
		},
		Commands: allCommands,
		Action: func(ctx *cli.Context) error {
			// Default action when no subcommand is specified
			if ctx.Bool("version") {
				printVersion(ctx)
				os.Exit(0)
				return nil
			}
			return cli.ShowAppHelp(ctx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err) //nolint:errcheck
		os.Exit(1)
	}
}

func printStdout(cCtx *cli.Context, format string, a ...any) { _ = "STUB: not implemented"; return }

//nolint:errcheck
