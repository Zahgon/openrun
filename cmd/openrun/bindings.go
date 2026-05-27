// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/openrundev/openrun/internal/types"
	"github.com/urfave/cli/v2"
)

const (
	SOURCE_FLAG       = "source"
	GRANT_FLAG        = "grant"
	ADD_GRANT_FLAG    = "add-grant"
	DELETE_GRANT_FLAG = "delete-grant"
	REAPPLY_ALL_FLAG  = "reapply-all"
	REAPPLY_ALL_ARG   = "reapplyAll"
)

func initBindingCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func bindingCreateCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func bindingUpdateCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func bindingDeleteCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func bindingGetCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func bindingShowAccountCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func bindingListCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func bindingRunCommand(commonFlags []cli.Flag, clientConfig *types.ClientConfig) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func printBindingList(cCtx *cli.Context, bindings []types.Binding, format string) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func formatMap(m map[string]string) string { _ = "STUB: not implemented"; return "" }
