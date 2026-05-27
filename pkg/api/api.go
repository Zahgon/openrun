// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	clserver "github.com/openrundev/openrun/internal/server"
	"github.com/openrundev/openrun/internal/types"
)

// ServerConfig is the configuration for the OpenRun Server
type ServerConfig struct {
	*types.ServerConfig
}

func NewServerConfig() (*ServerConfig, error) { _ = "STUB: not implemented"; return nil, nil }

// Server is the instance of the OpenRun Server
type Server struct {
	config *ServerConfig
	server *clserver.Server
}

// NewServer creates a new instance of the OpenRun Server
func NewServer(config *ServerConfig) (*Server, error) { _ = "STUB: not implemented"; return nil, nil }

// Start starts the OpenRun Server
func (s *Server) Start() error { _ = "STUB: not implemented"; return nil }

// Stop stops the OpenRun Server
func (s *Server) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
