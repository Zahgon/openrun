// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"context"
	"net/http"

	"github.com/openrundev/openrun/internal/types"
)

type DelegateRequest struct {
	ImageTag       string
	ContainerFile  string
	ContainerArgs  map[string]string
	RegistryConfig *types.RegistryConfig
}

func sendDelegateBuild(url string, data DelegateRequest, sourcePath string, builderAuthToken string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a pipe: writer feeds the HTTP request, reader is used as the body

// We need the content-type string for the request header *before* we close the writer.

// Build the multipart body in a goroutine so it streams

// Any error should close the pipe with that error

// Close the multipart writer first (writes the final boundary)

// Then close the pipe writer

// JSON part

// File part (streamed)

// nolint: errcheck

// After this, the deferred writer.Close + pw.Close run and finish the body.

// Create the request with the pipe reader as body

// Send it

// nolint: errcheck

// Read body for debugging

// DelegateHandler is the handler for the delegated build API
func DelegateHandler(r *http.Request, config *types.ServerConfig, logger *types.Logger) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Parse Content-Type to get boundary

// Stream-decode JSON for meta

// Ignore any extra fields; just drain them

func delegateBuild(ctx context.Context, logger *types.Logger, config *types.ServerConfig, data DelegateRequest, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

// nolint: errcheck

func pushToRemoteRegistry(ctx context.Context, logger *types.Logger, config *types.ServerConfig, imageTag string, registryConfig *types.RegistryConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// extractTarGzToTemp extracts the given .tar.gz file into a new temp directory.
// It returns the temp directory path on success.
func extractTarGzToTemp(tarGzPath string) (string, error) {
	_ = "STUB: not implemented"
	// Open the input file
	return "", nil
}

// nolint: errcheck

// Wrap in gzip reader

// nolint: errcheck

// Create a temp directory

// If we error later, clean up the temp dir

// done

// Ensure directory exists

// For now, ignore other entry types.

// Success: do not remove temp dir
