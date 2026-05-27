// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import "github.com/gorilla/sessions"

func sessionValueString(session *sessions.Session, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func mapValueString(values map[any]any, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func stateValueString(stateMap map[string]any, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func stateValueBool(stateMap map[string]any, key string) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func stateValueStringSlice(stateMap map[string]any, key string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func anyToStringSlice(raw any) ([]string, bool) { _ = "STUB: not implemented"; return nil, false }
