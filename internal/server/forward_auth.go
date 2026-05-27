// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"net/http"

	"github.com/openrundev/openrun/internal/types"
)

var hopByHopForwardAuthHeaders = map[string]struct{}{
	"Connection":          {},
	"Keep-Alive":          {},
	"Proxy-Authenticate":  {},
	"Proxy-Authorization": {},
	"Te":                  {},
	"Trailer":             {},
	"Transfer-Encoding":   {},
	"Upgrade":             {},
	"Content-Length":      {},
}

// forwardAuthMiddleware checks each request with the configured forward auth endpoint before it reaches the app handler.
func (s *Server) forwardAuthMiddleware(next http.Handler, forwardConfig *types.ForwardConfig) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

//nolint:errcheck

// newForwardAuthHTTPClient creates a forward auth client that returns redirects to the caller.
func newForwardAuthHTTPClient(config *types.ServerConfig) *http.Client {
	_ = "STUB: not implemented"
	return nil
}

// copyForwardAuthRequestHeaders copies configured original request headers to the auth request.
func copyForwardAuthRequestHeaders(authReq *http.Request, originalReq *http.Request, forwardHeaders []string) {
	_ = "STUB: not implemented"
	return
}

// copy all headers

// setForwardAuthHeaders sets the standard forward-auth request context headers for the auth endpoint.
func setForwardAuthHeaders(header http.Header, r *http.Request, config *types.ServerConfig) {
	_ = "STUB: not implemented"
	return
}

// setForwardAuthOpenRunHeaders adds OpenRun's authenticated identity and authorization context to the auth request.
func setForwardAuthOpenRunHeaders(header http.Header, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// deleteForwardAuthOpenRunHeaders removes client-supplied OpenRun identity headers before setting trusted values.
func deleteForwardAuthOpenRunHeaders(header http.Header) { _ = "STUB: not implemented"; return }

// copyForwardAuthResponseHeaders copies configured successful auth response headers onto the request sent to the app.
func copyForwardAuthResponseHeaders(requestHeader http.Header, responseHeader http.Header, copyResponseHeaders []string) {
	_ = "STUB: not implemented"
	return
}

// parseForwardAuthHeaderCopy parses a response header copy rule, including Caddy-style Source>Target renames.
func parseForwardAuthHeaderCopy(headerCopy string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// copyHeaders copies HTTP headers while omitting hop-by-hop headers that should not cross proxy boundaries.
func copyHeaders(dst http.Header, src http.Header) { _ = "STUB: not implemented"; return }

// shouldSkipForwardAuthHeader reports whether a header is hop-by-hop or otherwise unsafe to copy.
func shouldSkipForwardAuthHeader(headerName string) bool { _ = "STUB: not implemented"; return false }
