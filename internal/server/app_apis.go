// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"net/http"

	"github.com/openrundev/openrun/internal/app"
	"github.com/openrundev/openrun/internal/types"
)

func parseAppPath(inp string) (types.AppPathDomain, error) {
	_ = "STUB: not implemented"
	return *new(types.AppPathDomain), nil
}

func normalizePath(inp string) string {
	_ = "STUB: not implemented"
	// remove trailing slash
	return ""
}

func (s *Server) CreateApp(ctx context.Context, appPath string,
	approve, dryRun bool, appRequest *types.CreateAppRequest) (*types.AppCreateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) CreateAppTx(ctx context.Context, currentTx types.Transaction, appPath string,
	approve, dryRun bool, appRequest *types.CreateAppRequest, repoCache *RepoCache) (*types.AppCreateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If domain ends with a dot, append the default domain

// If source url has a hash, the part after the hash is the star base path

// Set the default for write access by staging and preview apps

// validated in createApp

func (s *Server) validateAppAuthnType(authStr string) error { _ = "STUB: not implemented"; return nil }

// saml auth, with or without rbac

func (s *Server) createApp(ctx context.Context, tx types.Transaction,
	appEntry *types.AppEntry, approve, dryRun bool, branch, commit, gitAuth string, applyInfo *types.CreateAppRequest, repoCache *RepoCache) (*types.AppCreateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make sure the source path is absolute

// Lowercase the ID, helps use the ID in container names

// Create the stage app entry if not dev

// Save the apply info in the app metadata (if called from apply context)

// Work on the stage app for prod apps, it will be promoted later

// Checkout the git repo locally and load into database

// App is loaded from disk (not git) and not in dev mode, load files into DB

// Create the in memory app object

// Persist the source url

// Persist the metadata so that any git info is saved

// Persist the settings

// Update the prod app metadata, promote from stage

// getAppHttpUrl returns the HTTP URL for accessing the app
func (s *Server) getAppHttpUrl(appEntry *types.AppEntry) string {
	_ = "STUB: not implemented"
	return ""
}

// getAppHttpsUrl returns the HTTPS URL for accessing the app
func (s *Server) getAppHttpsUrl(appEntry *types.AppEntry) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *Server) setupApp(ctx context.Context, appEntry *types.AppEntry, tx types.Transaction) (*app.App, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prod mode, use DB as source

// Dev mode, use local disk as source

func (s *Server) getAppBindings(ctx context.Context, inpTx types.Transaction, appEntry *types.AppEntry) ([]*types.Binding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) GetAppApi(ctx context.Context, appPath string) (*types.AppGetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (s *Server) GetAppEntry(ctx context.Context, tx types.Transaction, pathDomain types.AppPathDomain) (*types.AppEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) GetApp(ctx context.Context, pathDomain types.AppPathDomain, init bool) (*app.App, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// App not found in cache, get from DB

// Initialize the app

func (s *Server) DeleteApps(ctx context.Context, appPathGlob string, dryRun bool) (*types.AppDeleteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Remove from in memory app cache

func (s *Server) checkAuthModifiers(authTypeFull string) (string, *types.ForwardConfig, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (s *Server) authenticateAndServeApp(w http.ResponseWriter, r *http.Request, app *app.App) {
	_ = "STUB: not implemented"
	return
}

// no default auth type set, default to system admin user auth

// Check for auth modifiers which are used to implement forward_auth

// Remove the RBAC_AUTH_PREFIX rbac: prefix

// Cookie-based auth must start on the app's configured host. Otherwise a
// fallback-matched unknown host would receive the auth nonce/session cookie.

// No authentication required

// Use system admin user for authentication

// not using the actual user id, just a admin placeholder

// Use client certificate authentication

// Use SAML auth

// Already redirected to auth provider

// Use SSO auth

// Redirect to the auth provider if not logged in

// Already redirected to auth provider

// Create a new context with the user ID

// allow audit middleware to access the user id

// wrap the app with the csrf middleware

// Authentication successful, serve the app

func usesSessionCookieAuth(authType string) bool { _ = "STUB: not implemented"; return false }

func sameAppRequestDomain(requestDomain, appDomain string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Server) canonicalAuthRedirectURL(r *http.Request, appPathDomain types.AppPathDomain) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func isOpenRunCookieName(name string) bool { _ = "STUB: not implemented"; return false }

func isCookieWhitespace(b byte) bool { _ = "STUB: not implemented"; return false }

func stripOpenRunCookieHeader(cookieHeader string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func stripOpenRunCookies(r *http.Request) { _ = "STUB: not implemented"; return }

// verifyClientCerts verifies the client certificate, whether it is signed by one
// of the root CAs in the authName config
func (s *Server) verifyClientCerts(r *http.Request, authName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) MatchApp(hostHeader, matchPath string) (types.AppInfo, error) {
	_ = "STUB: not implemented"
	return *new(types.AppInfo), nil
}

// Request to unknown domain, match against default domain

// Host header does not match

// Do not match /_cl_stage to /

func (s *Server) CheckAppValid(domain, matchPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If /test is in use, do not allow /test/other

// If /test/other is in use, do not allow /test

func (s *Server) auditApp(ctx context.Context, tx types.Transaction, app *app.App, approve bool) (*types.ApproveResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) CompleteTransaction(ctx context.Context, tx types.Transaction, entries []types.AppPathDomain, dryRun bool, op string) error {
	_ = "STUB: not implemented"
	return nil
}

// Used when called in a context where the transaction is handled by the caller

// Update the in memory cache

func (s *Server) getStageApp(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry) (*types.AppEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const REPO_FOLDER_SEPERATOR = "//"

func parseGitUrl(sourceUrl string, usingSSH bool) (repo, folder string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// GitLab supports groups and subgroups, like gitlab.com/g16004341/g2/pr1/app1
// OpenRun requires such urls to be specified with // separator for the folder path
// like gitlab.com/g16004341/g2/pr1//app1

// Using git url format

// Use git url like git@github.com:openrundev/openrun.git

// Use git url like git@github.com:openrundev/openrun.git

type gitAuthEntry struct {
	user     string
	key      []byte
	password string
	usingSSH bool
}

// loadGitKey gets the git key from the config and loads the key from disk
func (s *Server) loadGitKey(gitAuth string) (*gitAuthEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default to non-SSH git url if no auth is specified

// https://github.com/src-d/go-git/issues/637, default user to "git"

func (s *Server) loadSourceFromGit(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry, branch, commit, gitAuth string, repoCache *RepoCache) error {
	_ = "STUB: not implemented"
	return nil
}

// Dev app from git, we need to point the app to the local checkout location

// App metadata points to the local checkout location

// Update the git info into the appEntry, the caller needs to persist it into the app metadata
// This function will persist it into the app_version metadata

// Walk the local directory and add all files to the database

// No previous version, start at 0

func (s *Server) loadSourceFromDisk(ctx context.Context, tx types.Transaction, appEntry *types.AppEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// No previous version, set to 0

// Walk the local directory and add all files to the database

func (s *Server) FilterApps(appappPathGlob string, includeInternal bool) ([]types.AppInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filter based on path spec. This is done on the main apps path only.

// Include staging and preview apps for prod apps

func (s *Server) GetApps(ctx context.Context, appPathGlob string, internal bool) ([]types.AppResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// staging app is at different version than prod app

func (s *Server) PreviewApp(ctx context.Context, mainAppPath, commitId string, approve, dryRun bool) (*types.AppPreviewResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Check if it already exists

// Checkout the git repo locally and load into database

// Create the in memory app object

// Persist the metadata so that any git info is saved

// Persist the settings

// Needs approval but not approved, do not create the preview app

// Clear the cache so that the new app is loaded next time
