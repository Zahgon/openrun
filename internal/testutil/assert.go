// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"testing"
)

func AssertEqualsString(tb testing.TB, msg, want, got string) { _ = "STUB: not implemented"; return }

// AssertStringMatch matches strings after removing extra spaces
func AssertStringMatch(tb testing.TB, msg, want, got string) { _ = "STUB: not implemented"; return }

func AssertEqualsInt(tb testing.TB, msg string, want, got int) { _ = "STUB: not implemented"; return }

func AssertEqualsBool(tb testing.TB, msg string, want, got bool) { _ = "STUB: not implemented"; return }

func AssertNoError(tb testing.TB, err error) { _ = "STUB: not implemented"; return }

func AssertErrorContains(tb testing.TB, err error, want string) { _ = "STUB: not implemented"; return }

func AssertEqualsError(tb testing.TB, msg string, got error, want error) {
	_ = "STUB: not implemented"
	return
}

func AssertStringContains(tb testing.TB, str string, want string) {
	_ = "STUB: not implemented"
	return
}
