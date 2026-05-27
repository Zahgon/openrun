// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// metricsEnabled is set true by Setup when metrics are configured. It is a
// finer-grained switch than Enabled(): callers that only care about counters
// (e.g. the SQL driver wrapper, the container manager wrapper) can avoid the
// per-call work entirely when metrics are off.
var metricsEnabled atomic.Bool

// MetricsEnabled reports whether metric instrumentation is currently active.
func MetricsEnabled() bool { _ = "STUB: not implemented"; return false }

var (
	dbInstrumentsOnce        sync.Once
	dbCallDuration           metric.Float64Histogram
	containerInstrumentsOnce sync.Once
	containerCallDuration    metric.Float64Histogram
	appInstrumentsOnce       sync.Once
	appRequest               metric.Int64Counter
	appResponse              metric.Int64Counter
	appProxyBytes            metric.Int64Counter
)

// resetMetricInstruments is called from Shutdown so that a subsequent Setup
// re-creates instruments against a fresh MeterProvider.
func resetMetricInstruments() { _ = "STUB: not implemented"; return }

func ensureDBInstruments() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func ensureContainerInstruments() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func ensureAppInstruments() bool { _ = "STUB: not implemented"; return false }

// RecordDBCall records the duration and outcome of a SQL driver call. It is a
// no-op when metrics are disabled.
func RecordDBCall(ctx context.Context, dbSystem, invoker, operation string, start time.Time, err error) {
	_ = "STUB: not implemented"
	return
}

// RecordContainerCall records the duration and outcome of a container manager
// call. It is a no-op when metrics are disabled.
func RecordContainerCall(ctx context.Context, kind, operation string, start time.Time, err error, extraAttrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

// RecordAppRequest records app-level request counters. It is a no-op when
// metrics are disabled.
func RecordAppRequest(ctx context.Context, method string, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

// RecordAppResponse records app-level HTTP status counters. It is a no-op when
// metrics are disabled.
func RecordAppResponse(ctx context.Context, status int, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

// RecordAppProxyBytes records app reverse-proxy byte counters. bytesIn is
// traffic received from the client, and bytesOut is traffic sent to the client.
func RecordAppProxyBytes(ctx context.Context, bytesIn, bytesOut uint64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func statusBucket(status int) string { _ = "STUB: not implemented"; return "" }

func metricAttrs(attrs []attribute.KeyValue, extra attribute.KeyValue) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func saturatingInt64(v uint64) int64 { _ = "STUB: not implemented"; return 0 }
