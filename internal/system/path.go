// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package system

// CleanRelativePath returns a slash-separated relative path that cannot escape
// its containing directory.
//
// For example, `sub\file.txt` is returned as `sub/file.txt`. Use this for
// logical paths such as fs.FS names, URL/static paths, metadata paths, and tar
// header names.
//
// Backslashes in the input are treated as path separators on every platform so
// that configs authored on Windows or Unix are validated identically. The
// trade-off is that a legitimate Unix filename containing a backslash will be
// re-interpreted as a multi-component path; callers that need to preserve
// such names should not route them through this function.
func CleanRelativePath(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Reject host-native absolute paths (e.g. "C:\\foo" on Windows) before any
// normalization, in case the platform parser disagrees with our slash
// rewrite below.

// CleanRelativeLocalPath returns a platform-local relative path that cannot
// escape its containing directory. On Windows, filepath.IsLocal additionally
// rejects reserved device names such as NUL, CON, COM1, etc.
//
// For example, `sub\file.txt` is returned as `sub/file.txt` on Unix and
// `sub\file.txt` on Windows. Use this when the result is passed to filepath,
// os, or host commands.
func CleanRelativeLocalPath(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// CleanFilename returns a single filename with no directory components.
//
// As with CleanRelativePath, backslashes are treated as path separators on
// every platform so that filenames are validated portably. On Windows the
// result is also checked against reserved device names (NUL, CON, COM1, ...).
func CleanFilename(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// filepath.IsLocal rejects Windows reserved device names like NUL or CON
// (this is a no-op on non-Windows platforms).

// PathInDir joins relPath under root after validating that relPath is local.
//
// Because CleanRelativeLocalPath enforces filepath.IsLocal, the resulting
// joined path is guaranteed to be lexically inside root. This is a lexical
// guarantee only: it does not resolve symlinks. For symlink-safe access use
// os.Root from the standard library.
func PathInDir(root, relPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// CleanAbsolutePath returns an absolute, cleaned local filesystem path. If the
// path exists, symlinks are resolved.
func CleanAbsolutePath(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// PathWithinDir reports whether targetPath is lexically inside root.
//
// This is a purely lexical check: symlinks under root that point outside root
// are not detected. For symlink-safe access prefer os.Root from the standard
// library.
func PathWithinDir(root, targetPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func hasWindowsDrivePrefix(name string) bool { _ = "STUB: not implemented"; return false }
