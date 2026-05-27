// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package system

import (
	"os"
	"os/exec"
)

// SetProcessGroup sets the process group flag for the command
func SetProcessGroup(cmd *exec.Cmd) { _ = "STUB: not implemented"; return }

// KillGroup kills the process group
func KillGroup(process *os.Process) error { _ = "STUB: not implemented"; return nil }
