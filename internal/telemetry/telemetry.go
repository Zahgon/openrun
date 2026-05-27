// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/openrundev/openrun/internal/types"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

const instrumentationName = "github.com/openrundev/openrun"

var (
	enabled         atomic.Bool
	pluginSpansOn   atomic.Bool
	emptyPropagator = propagation.NewCompositeTextMapPropagator()
)

// Providers owns the lifecycle of the SDK providers. It is always returned
// non-nil from Setup so callers do not need to nil-check before Shutdown.
type Providers struct {
	tracerProvider *sdktrace.TracerProvider
	meterProvider  *sdkmetric.MeterProvider
}

// Enabled reports whether telemetry is currently active.
func Enabled() bool { _ = "STUB: not implemented"; return false }

// PluginSpansEnabled reports whether per-plugin-call spans should be created.
// Plugin spans can be expensive in apps with high plugin call counts, so they
// are gated independently of the master telemetry switch.
func PluginSpansEnabled() bool { _ = "STUB: not implemented"; return false }

// Setup initializes OpenTelemetry providers based on the server config. It
// always returns a non-nil Providers value: when telemetry is disabled or when
// initialization fails, Shutdown becomes a no-op and the helper functions
// short-circuit through Enabled().
func Setup(ctx context.Context, config *types.ServerConfig, logger *types.Logger) (*Providers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shutdown flushes and closes the SDK providers. It is safe to call on a
// disabled or partially-initialized Providers value.
func (p *Providers) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Reset cached metric instruments so a subsequent Setup re-creates them
// against the new MeterProvider rather than reusing handles bound to the
// shut-down one.

// Reset globals so callers that cached a tracer don't keep emitting into a
// shut-down exporter (mostly relevant for tests that re-Setup in-process).

// Tracer returns the OpenRun tracer (a no-op when telemetry is disabled).
func Tracer() trace.Tracer { _ = "STUB: not implemented"; return *new(trace.Tracer) }

// Meter returns the OpenRun meter (a no-op when telemetry is disabled).
func Meter() metric.Meter { _ = "STUB: not implemented"; return *new(metric.Meter) }

// StartSpan starts a span with the given name and attributes. When telemetry
// is disabled it returns the input context and a no-op span.
func StartSpan(ctx context.Context, name string, attrs ...attribute.KeyValue) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

// RecordError annotates a span with an error if err is non-nil. Safe with a
// nil span.
func RecordError(span trace.Span, err error) { _ = "STUB: not implemented"; return }

// ServerHandlerOption configures WrapServerHandler.
type ServerHandlerOption struct {
	// Operation is the otelhttp base operation name.
	Operation string
	// Public marks the listener as untrusted. When true, incoming
	// trace-context headers are NOT extracted, so external clients cannot
	// inject parent spans into our trace.
	Public bool
	// ExtraSkipPaths lists URL path prefixes that should never be traced. The
	// internal health endpoint is always skipped.
	ExtraSkipPaths []string
	// TraceOnlyPrefixes, when set, limits server-level instrumentation to URL
	// path prefixes owned by OpenRun itself. Public app traffic should rely on
	// app-level spans where per-app redaction/skip policy is available.
	TraceOnlyPrefixes []string
}

// WrapServerHandler wraps an http.Handler with otelhttp instrumentation. When
// telemetry is disabled the handler is returned unchanged.
func WrapServerHandler(handler http.Handler, opt ServerHandlerOption) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Do not let untrusted clients become the parent of our spans.

// WrapTransport wraps an outbound RoundTripper with otelhttp instrumentation,
// which propagates the active trace context to the upstream and records a
// client span.
func WrapTransport(base http.RoundTripper) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

// AppIdentityAttributes returns the per-app identity attributes used on app
// traces and metrics.
func AppIdentityAttributes(app *types.AppEntry) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// AppAttributes returns the immutable per-app attribute set. Cache the result
// on the App and reuse across requests; do not call this on the hot path.
func AppAttributes(app *types.AppEntry) []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

// RequestAttributes returns per-request attributes that are safe to record.
// Notably, no client-supplied values (Host, URL path, query) are included
// here; those must be added by callers that have applied per-app redaction.
func RequestAttributes(r *http.Request) []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func traceExporterOptions(config *types.ServerConfig) []otlptracehttp.Option {
	_ = "STUB: not implemented"
	return nil
}

func metricExporterOptions(config *types.ServerConfig) []otlpmetrichttp.Option {
	_ = "STUB: not implemented"
	return nil
}

type otlpEndpoint struct {
	host     string
	fullURL  string
	insecure bool
}

func (e otlpEndpoint) configured() bool { _ = "STUB: not implemented"; return false }

func parseOTLPEndpoint(raw string) otlpEndpoint {
	_ = "STUB: not implemented"
	return *new(otlpEndpoint)
}

func newResource(ctx context.Context, config *types.ServerConfig) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serverSpanName(_ string, r *http.Request) string { _ = "STUB: not implemented"; return "" }

func routeClass(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func requestScheme(r *http.Request) string { _ = "STUB: not implemented"; return "" }
