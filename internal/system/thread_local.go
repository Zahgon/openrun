// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package system

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

func GetThreadLocalKey(thread *starlark.Thread, key string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetRequestUserId(thread *starlark.Thread) string { _ = "STUB: not implemented"; return "" }

func GetRequestContext(thread *starlark.Thread) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetRequestGroups(thread *starlark.Thread) []string { _ = "STUB: not implemented"; return nil }

func GetContextGroups(ctx context.Context) []string { _ = "STUB: not implemented"; return nil }

func GetContextValue(ctx context.Context, key types.ContextKey) string {
	_ = "STUB: not implemented"
	return ""
}

func GetContextUserId(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func GetContextUserSubject(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func GetContextUserEmail(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func GetContextRequestId(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func GetContextAppId(ctx context.Context) types.AppId {
	_ = "STUB: not implemented"
	return *new(types.AppId)
}

func GetCustomPerms(ctx context.Context) []string { _ = "STUB: not implemented"; return nil }

func IsAppRBACEnabled(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
