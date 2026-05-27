// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	"github.com/openrundev/openrun/internal/container"
	"github.com/openrundev/openrun/internal/types"
)

type staleContainerManager interface {
	ListOpenRunContainers(ctx context.Context) ([]container.Container, error)
	StopContainer(ctx context.Context, name container.ContainerName) error
}

func (s *Server) startStaleContainerCleanup() { _ = "STUB: not implemented"; return }

func (s *Server) staleContainerCleanupRunner() { _ = "STUB: not implemented"; return }

func (s *Server) cleanupStaleContainers(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanupStaleContainers(ctx context.Context, logger *types.Logger, manager staleContainerManager, active map[container.ContainerName]bool) error {
	_ = "STUB: not implemented"
	return nil
}
