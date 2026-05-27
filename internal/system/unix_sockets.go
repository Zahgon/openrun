// Code based on https://github.com/tv42/httpunix/blob/master/httpunix.go

package system

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"
)

// Scheme is the URL scheme used for HTTP over UNIX domain sockets.
const Scheme = "http+unix"

// Transport is a http.RoundTripper that connects to Unix domain
// sockets.
type Transport struct {
	// DialTimeout is deprecated. Use context instead.
	DialTimeout time.Duration
	// RequestTimeout is deprecated and has no effect.
	RequestTimeout time.Duration
	// ResponseHeaderTimeout is deprecated. Use context instead.
	ResponseHeaderTimeout time.Duration

	onceInit  sync.Once
	transport http.Transport

	mu sync.Mutex
	// map a URL "hostname" to a UNIX domain socket path
	loc map[string]string
}

func (t *Transport) initTransport() { _ = "STUB: not implemented"; return }

//nolint:staticcheck

func (t *Transport) getTransport() *http.Transport { _ = "STUB: not implemented"; return nil }

func (t *Transport) dialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (t *Transport) dialTLS(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// RegisterLocation registers an URL location and maps it to the given
// file system path.
//
// Calling RegisterLocation twice for the same location is a
// programmer error, and causes a panic.
func (t *Transport) RegisterLocation(loc string, path string) { _ = "STUB: not implemented"; return }

var _ http.RoundTripper = (*Transport)(nil)

// RoundTrip executes a single HTTP transaction. See
// net/http.RoundTripper.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get http.Transport to cooperate
