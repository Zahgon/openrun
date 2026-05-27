// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/openrundev/openrun/internal/types"
)

// OAuth and OIDC support using goth library. Standard OAuth flow, using
// KV-backed gorilla sessions so browser cookies only carry opaque session ids.
// Two unusual situations are handled:
// 1. The OAuth callback url is on domain a.com while the app is on b.com. The standard flow does not work since the callback api (on a.com) cannot
//  set cookies on b.com. This is handled by having a redirect api on b.com which sets the cookies
// 2. There could be multiple OpenRun server instances. The metadata database is used to save the session info after the user is authenticated.
//  This db entry is used in the redirect api to create the cookies, and then db entry is deleted.

// The flow is
// 1. At service startup, OAuthManager and all the providers are initialized
// 2. For apps using OAuth/OIDC, CheckAuth is called to check if the user is authenticated
// 3. CheckAuth verifies the session cookie to see if the user is authenticated. If yes, done
// 4. If the user is not authenticated, beginLogin function is called (API is currently on the app domain)
// 5. beginLogin creates a sessionid and nonce. Saves entry in DB with sessionid as key and state map as value
// 6. beginLogin creates a cookie with the nonce and redirect url. Calls the login API on the callback domain
// 7. Redirects to the OAuth provider's login page, with sessionid in state
// 8. Login API calls gothic.BeginAuthHandler to redirect to the OAuth provider's login page
// 9. OAuth provider's login page redirects to the callback api on the callback domain, with sessionid in state
// 10. Callback api validates the sessionid, and updates the state map in the DB with the user id and groups info
// 11. Callback api redirects to the redirect API on the app domain, again passing the sessionid in the state parameter
// 12. redirect API validates the passed sessionid, nonce from DB statemap against nonce from cookie,
// 13. redirect stores the authenticated session, with the user id and groups info and deletes the DB entry
// 14. Redirects back to original app url, which will again call CheckAuth and find the authenticated cookie

const (
	PROVIDER_NAME_DELIMITER = "_"
	AUTH_KEY                = "authenticated"
	USER_KEY                = "user" // email/userid/nickname (for git email/nickname/userid)
	USER_ID_KEY             = "user_id"
	USER_EMAIL_KEY          = "email"
	USER_NICKNAME_KEY       = "nickname"
	PROVIDER_NAME_KEY       = "provider_name"
	GROUPS_KEY              = "groups"
	SESSION_INDEX_KEY       = "session_index"
	NONCE_KEY               = "nonce"
	REDIRECT_URL            = "redirect"
)

// OAuthManager manages the OAuth providers and their configurations (also OIDC)
type OAuthManager struct {
	*types.Logger
	config          *types.ServerConfig
	cookieStore     sessions.Store
	providerConfigs map[string]*types.AuthConfig
	db              KVStore
}

type OAuthAuthInfo struct {
	UserId      string
	Groups      []string
	UserSubject string
	UserEmail   string
}

func NewOAuthManager(logger *types.Logger, config *types.ServerConfig, db KVStore) *OAuthManager {
	_ = "STUB: not implemented"
	return nil
}

func getProviderName(r *http.Request) (string, error) { _ = "STUB: not implemented"; return "", nil }

func genCookieName(provider string) string { _ = "STUB: not implemented"; return "" }

func (s *OAuthManager) Setup(sessionKey []byte, sessionBlockKey []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Set the store for gothic

// google supports hosted domain option

// azuread requires a resources array, setting nil for now

// auth0 requires a domain

// okta requires an org url

// openidConnect requires a discovery url

// Register the providers with goth

func (s *OAuthManager) RegisterRoutes(csrfMiddleware *http.CrossOriginProtection, mux *chi.Mux) {
	_ = "STUB: not implemented"
	return
}

// Start login process

// gothic.Logout(w, r) needs to be called on the callback domain, so not done
// Set user as not authenticated in session

// Set user as unauthenticated in session

func (s *OAuthManager) ValidateProviderName(provider string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *OAuthManager) ValidateAuthType(authType string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *OAuthManager) CheckAuth(w http.ResponseWriter, r *http.Request, appProvider string) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (s *OAuthManager) CheckAuthInfo(w http.ResponseWriter, r *http.Request, appProvider string) (OAuthAuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(OAuthAuthInfo), nil
}

// delete the session

//nolint:errcheck

// Check if provider name matches the one in the session

// do the OAuth login flow

func (s *OAuthManager) beginLogin(w http.ResponseWriter, r *http.Request, providerName, redirectUrl string) {
	_ = "STUB: not implemented"
	return
}

// Store the state map in the database with the session id as the key

// Save a cookie with the nonce (this is on the app domain, not the callback domain)

// The state is the session id, encoded in base64

func (s *OAuthManager) authCallback(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// For git providers, prefer nickname over userid as it is more meaningful

// Get groups from user.RawData

// Update the state map, set to authenticated and add the user id and groups

// Redirect to the auth/redirect url on the original app domain, so that the required cookie can be set on the app domain

func (s *OAuthManager) redirect(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// delete the session if there is an error

// Get the state map, delete the entry from database, validate state, set the session values
// in the cookie and then redirect to original url

// Update the session cookie to authenticated, with the new values
