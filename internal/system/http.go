// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package system

import (
	"net"
	"net/http"
	"net/url"
)

const (
	ApplicationJson        = "application/json"
	OpenRunServiceLocation = "openrun"
)

type HttpClient struct {
	client    *http.Client
	serverUri string
	user      string
	password  string
}

// NewHttpClient creates a new HttpClient instance
func NewHttpClient(serverUri, user, password string, skipCertCheck bool) *HttpClient {
	_ = "STUB: not implemented"
	return nil
}

// Change to OPENRUN_HOME directory, helps avoid length limit on UDS file (around 104 chars)

// use relative path

// Using unix domain sockets

func (h *HttpClient) Get(url string, params url.Values, output any) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpClient) Post(url string, params url.Values, input any, output any) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpClient) Put(url string, params url.Values, input any, output any) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpClient) Delete(url string, params url.Values, output any) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HttpClient) request(method, apiPath string, params url.Values, input any, output any) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func MapServerHost(host string) string { _ = "STUB: not implemented"; return "" }

// GetRequestScheme returns "https" if the request was received over TLS locally,
// or if the direct peer is listed in trustedProxies and set X-Forwarded-Proto: https.
// Otherwise it returns "http". Only the first value of X-Forwarded-Proto is honored
// and only when the direct peer is a trusted proxy.
func GetRequestScheme(r *http.Request, trustedProxies []string) string {
	_ = "STUB: not implemented"
	return ""
}

// When a request traverses multiple proxies the header can be a
// comma-separated list like "https, http"; the leftmost value is
// the client-facing scheme. SplitN with n=2 avoids scanning the
// full string when there are many hops.

// IsOrigRequestHTTPS reports whether the request is (or originally came in as) HTTPS,
// honoring X-Forwarded-Proto only when the direct peer is a trusted proxy.
func IsOrigRequestHTTPS(r *http.Request, trustedProxies []string) bool {
	_ = "STUB: not implemented"
	return false
}

func GetRequestUrl(r *http.Request, trustedProxies []string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetHostname returns the hostname portion of an HTTP host header, handling
// hostnames, IPv4 addresses, and bracketed or bare IPv6 literals.
func GetHostname(host string) string { _ = "STUB: not implemented"; return "" }

// GetClientIP returns the caller IP, honoring forwarding headers only when the
// direct peer is explicitly configured as a trusted proxy.
func GetClientIP(r *http.Request, trustedProxies []string) string {
	_ = "STUB: not implemented"
	return ""
}

func forwardedClientIP(values []string, trustedProxies []string) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

func isTrustedProxy(ip net.IP, trustedProxies []string) bool {
	_ = "STUB: not implemented"
	return false
}

func parseIPValue(value string) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }
