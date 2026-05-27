// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v2"
)

const (
	FORMAT_TABLE        = "table"
	FORMAT_BASIC        = "basic"
	FORMAT_JSON         = "json"
	FORMAT_JSONL        = "jsonl"
	FORMAT_JSONL_PRETTY = "jsonl_pretty"
	FORMAT_CSV          = "csv"
)
const (
	//Terminal colors
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
)

func newStringFlag(name, alias, usage, value string) *cli.StringFlag {
	_ = "STUB: not implemented"
	return nil
}

func newIntFlag(name, alias, usage string, value int) *cli.IntFlag {
	_ = "STUB: not implemented"
	return nil
}

func newBoolFlag(name, alias, usage string, value bool) *cli.BoolFlag {
	_ = "STUB: not implemented"
	return nil
}

func validateNoFlagLikeValues(flagName string, valueName string, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

// makeAbsolute converts a relative path to an absolute path.
// This needs to be called in the client before the call to system.NewHttpClient
// since that changes the cwd to $OPENRUN_HOME
func makeAbsolute(sourceUrl string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Convert to absolute path so that server can find it
