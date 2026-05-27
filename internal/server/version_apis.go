// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
)

func (s *Server) VersionList(ctx context.Context, mainAppPath string) (*types.AppVersionListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) VersionFiles(ctx context.Context, mainAppPath, version string) (*types.AppVersionFilesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) VersionSwitch(ctx context.Context, mainAppPath string, dryRun bool, version string) (*types.AppVersionSwitchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Find the next valid version which is present

// Find the previous valid version which is present

// Don't commit the transaction if its a dry run
