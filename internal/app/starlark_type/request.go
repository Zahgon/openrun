// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package starlark_type

import (
	"net/http"
	"net/url"

	"go.starlark.net/starlark"
)

// Request is a starlark.Value that represents an HTTP request. A Request is created from the Go http.Request
// and passed to the starlark handler function as it only argument. The Data field is updated with the handler's
// response and then the template evaluation is done with the same Request
type Request struct {
	AppName        string
	AppPath        string
	AppUrl         string
	PagePath       string
	PageUrl        string
	Method         string
	IsDev          bool
	IsPartial      bool
	PushEvents     bool
	HtmxVersion    string
	Headers        http.Header
	RemoteIP       string
	UrlParams      map[string]string
	Form           url.Values
	Query          url.Values
	PostForm       url.Values
	UserId         string
	UserSubject    string
	UserEmail      string
	CustomPerms    []string
	AppRBACEnabled bool
	Data           any
}

func (r Request) Attr(name string) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (r Request) AttrNames() []string { _ = "STUB: not implemented"; return nil }

func (r Request) String() string { _ = "STUB: not implemented"; return "" }

func (r Request) Type() string { _ = "STUB: not implemented"; return "" }

func (r Request) Freeze() { _ = "STUB: not implemented"; return }

func (r Request) Truth() starlark.Bool { _ = "STUB: not implemented"; return *new(starlark.Bool) }

func (r Request) Hash() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

var _ starlark.Value = (*Request)(nil)
