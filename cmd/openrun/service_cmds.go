// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/openrundev/openrun/internal/types"
	"github.com/urfave/cli/v2"
)

const (
	SET_DEFAULT_FLAG = "set-default"
	IS_DEFAULT_FLAG  = "is-default"
	CONFIG_FLAG      = "config"
	STAGING_FLAG     = "staging"
)

func initServiceCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

// parseServiceID parses a service id of the form <service_type>[/<service_name>].
// If name is omitted, it defaults to the service type.
func parseServiceID(id string) (serviceType, name string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func parseConfigEntries(entries []string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serviceCreateCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func serviceUpdateCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

// Fetch the existing service to merge changes onto

func serviceDeleteCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func serviceListCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func printServiceList(cCtx *cli.Context, services []types.Service, format string) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func formatConfig(config map[string]string) string { _ = "STUB: not implemented"; return "" }
