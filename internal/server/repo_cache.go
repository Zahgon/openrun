// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"github.com/go-git/go-git/v5/plumbing/transport" // for AuthMethod
)

type Repo struct {
	url    string
	branch string
	commit string
	auth   string
}

type CacheDir struct {
	dir           string
	commitMessage string
	hash          string
}

type RepoCache struct {
	server   *Server
	rootDir  string
	cache    map[Repo]CacheDir
	shaCache map[Repo]string // Cache for commit hashes
}

func NewRepoCache(server *Server) (*RepoCache, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *RepoCache) GetSha(sourceUrl, branch, gitAuth string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Figure on which repo to clone

// Check if we have the commit in cache

func (r *RepoCache) createAuthMethod(gitAuth string) (transport.AuthMethod, error) {
	_ = "STUB: not implemented"
	return *new(transport.AuthMethod), nil
}

// SSH auth

// HTTP auth, either basic or using Personal Access Token

func latestCommitSHA(repoURL, branch string, auth transport.AuthMethod) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// e.g. "refs/heads/main"

func (r *RepoCache) CheckoutRepo(sourceUrl, branch, commit, gitAuth string, isDev bool) (string, string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil
}

// Figure on which repo to clone

// Don't fetch tags, to speed up checkout

// No commit id specified, checkout specified branch

// We don't have a previous dev checkout for this repo, create a new one

// Configure the repo to Clone

// Checkout specified hash

/* Sparse checkout seems to not be reliable with go-git
if folder != "" {
	options.SparseCheckoutDirectories = []string{folder}
}
*/

// Save the repo in cache

func getUnusedRepoPath(targetDir, repoName string) string { _ = "STUB: not implemented"; return "" }

func (r *RepoCache) Cleanup() { _ = "STUB: not implemented"; return }

//nolint:errcheck
