// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"time"
)

// Binding is a binding entry in the metadata database
// A binding is a link between a service and a source service
type Binding struct {
	Id             string          `json:"id"`           // the id of the binding
	Path           string          `json:"path"`         // the path of the binding
	Source         string          `json:"source"`       // service id, or the base binding path
	ServiceType    string          `json:"service_type"` // the type of the service
	ServiceName    string          `json:"service_name"` // the name of the service
	DerivedFrom    string          `json:"derived_from"` // the base binding path this is derived from
	StagedMetadata BindingMetadata `json:"staged_metadata"`
	Metadata       BindingMetadata `json:"metadata"`
	CreateTime     time.Time       `json:"create_time"`
	UpdateTime     time.Time       `json:"update_time"`
}

type BindingMetadata struct {
	Grants        []string          `json:"grants"`
	GrantsApplied []BindingGrant    `json:"grants_applied"`
	Config        map[string]string `json:"config"`
	Account       map[string]string `json:"account,omitempty"`
	ApplyInfo     []byte            `json:"apply_info"`
}

type GrantType string

const (
	GrantTypeRead   GrantType = "READ"
	GrantTypeCreate GrantType = "CREATE"
	GrantTypeFull   GrantType = "FULL"
)

const (
	GrantTargetAll = "*"
)

func ParseGrant(grant string, supportedGrantTypes []GrantType) (BindingGrant, error) {
	_ = "STUB: not implemented"
	return *new(BindingGrant), nil
}

type BindingGrant struct {
	GrantType   GrantType `json:"grant_type"`
	GrantTarget string    `json:"grant_target"`
}

func (g BindingGrant) String() string { _ = "STUB: not implemented"; return "" }
