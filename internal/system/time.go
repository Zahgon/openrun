// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package system

import (
	"time"
)

// HumanDuration returns a human readable duration string
func HumanDuration(d time.Duration) string { _ = "STUB: not implemented"; return "" }

// Round to whole seconds so we don't show sub-second noise.
