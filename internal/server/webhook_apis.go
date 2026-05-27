// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
)

func (s *Server) getServerUri() string { _ = "STUB: not implemented"; return "" }

func (s *Server) TokenList(ctx context.Context, appPath string) (*types.TokenListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) TokenCreate(ctx context.Context, appPath string, webhookType types.WebhookType, dryRun bool) (*types.TokenCreateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Persist the settings

func (s *Server) TokenDelete(ctx context.Context, appPath string, webhookType types.WebhookType, dryRun bool) (*types.TokenDeleteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Persist the settings
