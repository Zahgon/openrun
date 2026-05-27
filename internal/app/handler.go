// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"sync"

	"github.com/openrundev/openrun/internal/app/starlark_type"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var (
	CONTENT_TYPE_JSON = []string{"application/json"}
	CONTENT_TYPE_TEXT = []string{"text/plain"}

	SERVER_NAME       = []string{"OpenRun"}
	VARY_HEADER_VALUE = []string{"HX-Request"}
)

func (a *App) earlyHints(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func (a *App) getRequestUrl(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func defaultPortForScheme(scheme string) string { _ = "STUB: not implemented"; return "" }

func sameOriginURL(refURL *url.URL, requestURL *url.URL) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *App) validatedRefererRedirect(r *http.Request, referrer string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// pooled holds one encoder and its buffer.
type pooled struct {
	enc *json.Encoder
	buf *bytes.Buffer
}

var encoderPool = sync.Pool{
	New: func() interface{} {
		buf := bytes.NewBuffer(make([]byte, 0, 512))
		return &pooled{
			enc: json.NewEncoder(buf),
			buf: buf,
		}
	},
}

func (a *App) createHandlerFunc(fullHtml, fragment string, handler starlark.Callable, rtype string) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Save the request context in the starlark thread local

//nolint:staticcheck

//nolint:staticcheck
// Prod mode, for a GET request from newer browsers on a top level HTML page, send http early hints

//nolint:errcheck // ignore error if no form data is passed

// no handler means empty Data map is passed into template

// Check for any deferred cleanups

// Audit event was set, insert it

//nolint:errcheck

// Call the handler function

// handle as if the handler had returned an error

// Iterate through the CallFrame stack for debugging information

// No err handler defined, abort

// error handler is defined, call it

//nolint:errcheck

// error handler itself failed

// Iterate through the CallFrame stack for debugging information

// response type struct returned by handler Instead of template defined in
// the route, use the template specified in the response

// Response from handler, or if handler failed, response from error_handler if defined

//nolint:staticcheck
// If the route type is JSON, then return the handler response as JSON

// If the route type is TEXT, then return the handler response as text

// If block is defined, and this is a non-GET request, then redirect to the referrer page
// This handles the Post/Redirect/Get pattern required if HTMX is disabled

func (a *App) callStarlarkHandler(r *http.Request, thread *starlark.Thread, handler starlark.Callable, args starlark.Tuple) (starlark.Value, error) {
	_ = "STUB: not implemented"
	return *new(starlark.Value), nil
}

func (a *App) executeTemplateTraced(r *http.Request, w http.ResponseWriter, fullHtml, fragment string, data any) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *App) handleResponse(retStruct *starlarkstruct.Struct, r *http.Request, w http.ResponseWriter, requestData starlark_type.Request, rtype string, deferredCleanup func() error) (bool, error) {
	_ = "STUB: not implemented"
	// Handle ace.redirect type struct returned by handler
	return false, nil
}

// starlark Type() is not implemented for structs, so we can't check the type
// Looked at the mandatory properties to decide on type for now

// Redirect type struct returned by handler

// Handle ace.response type struct returned by handler

// Default to the type set at the route level

// If the route type is JSON, then return the handler response as JSON

// If the route type is TEXT, then return the handler response as plain text

func (a *App) getRemoteIP(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func (a *App) handleStreamResponse(w http.ResponseWriter, r *http.Request, rtype string, fragment string, streamResponse map[string]any) {
	_ = "STUB: not implemented"
	// Stream the response to the client
	return
}

//nolint:staticcheck

//nolint:errcheck
