// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	"github.com/openrundev/openrun/internal/bindings"
	"github.com/openrundev/openrun/internal/types"
)

func newPrefixedId(prefix string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s *Server) validateStagingService(ctx context.Context, tx types.Transaction, service *types.Service) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) CreateService(ctx context.Context, service *types.Service, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

//nolint:errcheck

// First service of this type automatically becomes the default

// Clear any existing default for this service_type

func (s *Server) UpdateService(ctx context.Context, service *types.Service, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

// Clear default flag on any other service of this type

func (s *Server) DeleteService(ctx context.Context, name, serviceType string, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func (s *Server) ListServices(ctx context.Context, serviceType, name string) ([]*types.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) CreateBinding(ctx context.Context, binding *types.Binding, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

// Reference another binding by path - derived binding

// Reject multi-level nesting. A derived binding must be derived from a
// base binding (one whose Source points at a service, not at another
// binding). Allowing derived-of-derived would make ALTER DEFAULT
// PRIVILEGES reference the wrong creator role

// Base binding

// Reference a service by type alone

// Reference a service by type and name

// Not dry run, generate the account info
// Generate the staging account info, either against the staging service if set or against the main service

// Generate the production account info
// This runs as a separate transaction, since it might not be against same database as the metadata database

func (s *Server) getServiceBinding(ctx context.Context, service *types.Service, binding *types.Binding) (bindings.ServiceBinding, error) {
	_ = "STUB: not implemented"
	return *new(bindings.ServiceBinding), nil
}

func (s *Server) generateAccount(ctx context.Context, dryRun bool, service *types.Service, binding *types.Binding, derivedFrom *types.Binding, isStaging, reapplyAll bool) (map[string]string, []types.BindingGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//nolint:errcheck

//nolint:errcheck

func (s *Server) applyBindingGrants(ctx context.Context, dryRun bool, service *types.Service,
	binding *types.Binding, derivedFrom *types.Binding, isStaging bool, reapplyAll bool) ([]types.BindingGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

func normalizeGrantForStorage(grant string) string { _ = "STUB: not implemented"; return "" }

func normalizeGrantList(grants []string) []string { _ = "STUB: not implemented"; return nil }

func mergeGrantUpdates(current, addGrants, deleteGrants []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) UpdateBinding(ctx context.Context, updateRequest types.UpdateBindingRequest, dryRun, promote, reapplyAll bool) (*types.Binding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) DeleteBinding(ctx context.Context, path string, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func (s *Server) GetBinding(ctx context.Context, path string) (*types.Binding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// GetBindingWithAccount gets the binding with the account info un-redacted.
func (s *Server) GetBindingWithAccount(ctx context.Context, tx types.Transaction, path string) (*types.Binding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) ListBindings(ctx context.Context, source string) ([]*types.Binding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func redactBindingAccount(binding *types.Binding) *types.Binding {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) GetBindingAccount(ctx context.Context, path string, useStaging bool) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) RunBindingCommand(ctx context.Context, bindingName string, useStaging bool, command string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck
