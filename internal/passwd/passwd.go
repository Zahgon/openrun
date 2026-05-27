// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package passwd

const (
	PASSWORD_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#%^&*()_-+=<>?/|"
	BCRYPT_COST    = 10
)

// GenerateRandomPassword generates a random password
func generateRandString(length int, charsAllowed string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GeneratePassword generates a random password
func GeneratePassword() (string, error) { _ = "STUB: not implemented"; return "", nil }

func GenerateSessionNonce() (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

func GenerateRandomKey(length int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
