// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func NewLogger(config *LogConfig) *Logger { _ = "STUB: not implemented"; return nil }

func RollingFileLogger(config *LogConfig, logType string) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}
