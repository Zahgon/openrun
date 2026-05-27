// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

// Based on code from https://github.com/qri-io/starlib/blob/master/http/http.go

package plugins

import (
	"context"
	"net/http"

	"github.com/openrundev/openrun/internal/app"
	"github.com/openrundev/openrun/internal/plugin"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

// AsString unquotes a starlark string value
func AsString(x starlark.Value) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Encodings for form data.
//
// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST
const (
	formEncodingMultipart = "multipart/form-data"
	formEncodingURL       = "application/x-www-form-urlencoded"
	defaultTimeoutSeconds = 300
)

func init() {
	h := &httpPlugin{}
	pluginFuncs := []plugin.PluginFunc{
		app.CreatePluginApi(h.Get, app.READ),
		app.CreatePluginApi(h.Head, app.READ),
		app.CreatePluginApi(h.Options, app.READ),
		app.CreatePluginApi(h.Post, app.WRITE),
		app.CreatePluginApi(h.Put, app.WRITE),
		app.CreatePluginApi(h.Delete, app.WRITE),
		app.CreatePluginApi(h.Patch, app.WRITE),
	}
	app.RegisterPlugin("http", NewHttpPlugin, pluginFuncs)
}

type httpPlugin struct {
	client *http.Client
}

func NewHttpPlugin(pluginContext *types.PluginContext) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (h *httpPlugin) Get(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (h *httpPlugin) Head(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (h *httpPlugin) Options(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (h *httpPlugin) Post(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (h *httpPlugin) Put(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (h *httpPlugin) Delete(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (h *httpPlugin) Patch(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// reqMethod is a factory function for generating starlark builtin functions for different http request methods
func (h *httpPlugin) reqMethod(method string) func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return nil
}

// If the url starts with the container url, we need to replace it with the container proxy url

// 1xx and 3xx are also failed by default
//nolint:errcheck

func setQueryParams(rawurl *string, params *starlark.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func setBasicAuth(req *http.Request, auth starlark.Tuple) error {
	_ = "STUB: not implemented"
	return nil
}

func getKeyAsString(dict *starlark.Dict, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func setSignAuth(req *http.Request, auth *starlark.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func createSLAuthHeader(req *http.Request, userId, apiKey string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If IPv6 address, unescape the % chars if present and add square brackets

// Referer header is not include in signature

// Add date header if not already present

func setHeaders(req *http.Request, headers *starlark.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func setBody(req *http.Request, body starlark.String, formData *starlark.Dict, formEncoding starlark.String, jsondata starlark.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// Specifying the Content-Length ensures that https://go.dev/src/net/http/transfer.go doesnt specify Transfer-Encoding: chunked which is not supported by some endpoints.
// This is required when using ioutil.NopCloser method for the request body (see ShouldSendChunkedRequestBody() in the library mentioned above).

//nolint:errcheck

// Response represents an HTTP response, wrapping a go http.Response with
// starlark methods
type Response struct {
	http.Response
	thread     *starlark.Thread
	cleanupKey string
	cancel     context.CancelFunc
}

func (r *Response) cleanupBody(clearCleanup bool) error { _ = "STUB: not implemented"; return nil }

// Struct turns a response into a *starlark.Struct
func (r *Response) Struct() *starlarkstruct.Struct { _ = "STUB: not implemented"; return nil }

// HeadersDict flops
func (r *Response) HeadersDict() *starlark.Dict { _ = "STUB: not implemented"; return nil }

// Text returns the raw data as a string
func (r *Response) Text(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// reset reader to allow multiple calls

// JSON attempts to parse the response body as JSON
func (r *Response) JSON(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

// reset reader to allow multiple calls
