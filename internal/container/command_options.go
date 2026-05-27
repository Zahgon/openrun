// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package container

const regexAllowedContainerArgPrefix = "regex:"

// CommandOptionArgs converts parsed container options into CLI args.
// Built-in OpenRun options are parsed explicitly. Any remaining Docker/Podman
// flags must be listed in allowedContainerArgs before they are emitted.
func CommandOptionArgs(options CommandOptions, allowedContainerArgs map[string]string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RedactEnvArgs(args []string) []string { _ = "STUB: not implemented"; return nil }

func redactEnvAssignment(arg string) string { _ = "STUB: not implemented"; return "" }

func commandOtherOptionArgs(options map[string]any, allowedContainerArgs map[string]string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateAllowedContainerArg(key, value string, allowedContainerArgs map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}
