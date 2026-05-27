package app

import (
	"bufio"
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type bucket struct {
	sec  int64
	sent uint64
	recv uint64
}

type ByteWindow struct {
	mu      sync.Mutex
	buckets []bucket // len = windowSeconds
	window  int      // seconds
	attrs   []attribute.KeyValue
}

func NewByteWindow(windowSeconds int, attrs ...attribute.KeyValue) *ByteWindow {
	_ = "STUB: not implemented"
	return nil
}

func (bw *ByteWindow) add(ctx context.Context, now time.Time, sent, recv uint64) {
	_ = "STUB: not implemented"
	return
}

// this slot is stale; reset

func (bw *ByteWindow) Totals() (sent, recv uint64) { _ = "STUB: not implemented"; return 0, 0 }

type countingReadCloser struct {
	rc  io.ReadCloser
	bw  *ByteWindow
	ctx context.Context
}

func (c *countingReadCloser) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *countingReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

type countingResponseWriter struct {
	http.ResponseWriter
	bw  *ByteWindow
	ctx context.Context
}

func (w *countingResponseWriter) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *countingResponseWriter) Unwrap() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *

	// Support Flush for streaming/SSE.
	new(http.ResponseWriter)
}

func (w *countingResponseWriter) Flush() { _ = "STUB: not implemented"; return }

// Implement Hijacker so reverse proxy can upgrade to WS,
// and we can wrap the net.Conn to count both directions.
func (w *countingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

type countingConn struct {
	net.Conn
	bw  *ByteWindow
	ctx context.Context
}

func (c *countingConn) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// client -> proxy

func (c *countingConn) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// proxy -> client

// Tracker is a reverse proxy with byte count tracking
type Tracker struct {
	bw    *ByteWindow
	proxy *httputil.ReverseProxy
}

func NewTracker(proxy *httputil.ReverseProxy, windowSeconds int, attrs ...attribute.KeyValue) *Tracker {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tracker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// Count request bytes (body) for non-upgraded requests.
	return
}

// Accessor to read the rolling totals.
func (t *Tracker) GetRollingTotals() (sent, recv uint64) { _ = "STUB: not implemented"; return 0, 0 }
