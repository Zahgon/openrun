// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"context"
)

const defaultTestContainerCommand = "docker"

// StartPostgresContainer starts a Postgres container through the local container CLI.
func StartPostgresContainer(ctx context.Context, image, database, username, password string) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func randomContainerName() (string, error) { _ = "STUB: not implemented"; return "", nil }

func containerCleanup(containerCommand, containerID string) func() {
	_ = "STUB: not implemented"
	return nil
}

func waitForPublishedPort(ctx context.Context, containerCommand, containerID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func inspectPublishedPort(ctx context.Context, containerCommand, containerID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
