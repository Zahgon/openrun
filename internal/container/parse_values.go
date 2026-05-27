// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"regexp"
)

var (
	reIntOnly     = regexp.MustCompile(`^\d+$`)
	reDockerLike  = regexp.MustCompile(`^\d+(\.\d+)?\s*[bkmgte]b?\s*$`) // e.g. 512m, 1g, 1gb, 0.5g (case-insensitive handled below)
	KNOWN_OPTIONS = []string{"cpus", "memory", "min_replicas", "max_replicas"}
)

// BytesString parses s and returns bytes as a base-10 integer string.
//
// Rules:
// 1) If already an integer string (bytes), return as-is.
// 2) If docker-like (e.g., 512m, 1g), parse via Docker and return bytes.
// 3) Otherwise parse as k8s Quantity (e.g., 512Mi, 1Gi, 500M) and return bytes.
func BytesString(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// already integer bytes

// docker-like,  avoids treating "512m" as k8s milli-bytes

// fall through to try k8s (in case it wasn't really docker)

// k8s quantity -> bytes

// CPUString converts CPU from either docker-like ("0.5", "2") or k8s-like ("500m", "1")
// into a string that "makes sense" for the target.
//   - targetIsDocker=true  => return cores as decimal string (e.g. "0.5", "2")
//   - targetIsDocker=false => return millicores as integer string (e.g. "500", "2000")
//
// Notes:
//   - Bare "1" is treated as 1 core (not 1 millicore).
//   - For millicores input, use "m" suffix: "500m".
func CPUString(s string, targetIsDocker bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// integer millicores (ceil for fractional smaller than 1m)

func formatCoresFromMilli(m int64) string { _ = "STUB: not implemented"; return "" }

// exact 3-decimal formatting, then trim trailing zeros
