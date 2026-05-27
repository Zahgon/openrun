// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package action

import (
	"context"
	"embed"
	"html/template"
	"net/http"

	"github.com/benbjohnson/hashfs"
	"github.com/go-chi/chi/v5"
	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/app/apptype"
	"github.com/openrundev/openrun/internal/rbac"
	"github.com/openrundev/openrun/internal/types"
	"go.starlark.net/starlark"
)

//go:embed *.go.html astatic/*
var embedHtml embed.FS
var embedFS = hashfs.NewFS(embedHtml)

const (
	defaultMaxRequestBodyBytes int64 = 32 << 20
	multipartMaxMemoryBytes    int64 = 10 << 20
)

type ActionLink struct {
	Name       string
	Path       string
	Permits    []string
	Authorized bool
}

// Action represents a single action that is exposed by the App. Actions
// provide a way to trigger app operations, with an auto-generated form UI
// and an API interface
type Action struct {
	*types.Logger
	isDev               bool
	name                string
	description         string
	appPath             string
	run                 starlark.Callable
	suggest             starlark.Callable
	params              []apptype.AppParam
	paramValuesStr      map[string]string
	paramDict           starlark.StringDict
	actionTemplate      *template.Template
	pagePath            string
	AppTemplate         *template.Template
	StyleType           types.StyleType
	LightTheme          string
	DarkTheme           string
	containerProxyUrl   string
	hidden              map[string]bool // params which are not shown in the UI
	Links               []ActionLink    // links to other actions
	showValidate        bool
	auditInsert         func(*types.AuditEvent) error
	containerHandler    any // Container manager, if available, used to run commands in the container
	esmLibs             []types.JSLibrary
	appPathDomain       types.AppPathDomain
	serverConfig        *types.ServerConfig
	maxRequestBodyBytes int64
	permit              []string
	rbacApi             rbac.RBACAPI
}

// NewAction creates a new action
func NewAction(logger *types.Logger, sourceFS *appfs.SourceFs, isDev bool, name, description, apath string, run, suggest starlark.Callable,
	params []apptype.AppParam, paramValuesStr map[string]string, paramDict starlark.StringDict,
	appPath string, styleType types.StyleType, containerProxyUrl string, hidden []string, showValidate bool,
	auditInsert func(*types.AuditEvent) error, containerManager any, jsLibs []types.JSLibrary, appPathDomain types.AppPathDomain,
	serverConfig *types.ServerConfig, actionConfig types.ActionConfig, permit []string, rbacApi rbac.RBACAPI) (*Action, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Links, AppTemplate and Theme names are initialized later

func (a *Action) GetLink() ActionLink { _ = "STUB: not implemented"; return *new(ActionLink) }

// GetEmbeddedTemplates returns the embedded templates files
func GetEmbeddedTemplates() (map[string][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Action) BuildRouter() (*chi.Mux, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Action) runAction(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a *Action) suggestAction(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a *Action) validateAction(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a *Action) authorizeAction(w http.ResponseWriter, r *http.Request) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Action) execAction(w http.ResponseWriter, r *http.Request, isSuggest, isValidate bool, op string) {
	_ = "STUB: not implemented"
	return
}

// Audit event was set in handler, insert it

// Save the request context in the starlark thread local
// Same code as createHandlerFunc

// Check for any deferred cleanups

//nolint:errcheck

// Make a copy of the app level param dict

// Update args with submitted form values

//nolint:errcheck

//nolint:errcheck

// Write contents of uploaded file to destFile

// Not file upload, regular param

// Form does not submit unchecked checkboxes, set to false

// Call the handler function

// handle as if the handler had returned an error

// Iterate through the CallFrame stack for debugging information

// err handler is not supported for actions

// Not a result struct

// Set the push URL for HTMX

// Render the result message

// Render the param error messages, using HTMX OOB

// "" error messages have to be sent to overwrite previous values in form UI

// No need to render the results

func writeRequestParseError(w http.ResponseWriter, err error, maxRequestBodyBytes int64) {
	_ = "STUB: not implemented"
	return
}

func uploadedFilePath(tempDir, filename string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *Action) renderResults(w http.ResponseWriter, report string, valuesMap []map[string]any, valuesStr []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Custom template being used for the results
// Wrap the template output in a div with hx-swap-oob

func (a *Action) renderResultsAuto(w http.ResponseWriter, valuesMap []map[string]any, valuesStr []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Action) renderResultsText(w http.ResponseWriter, valuesStr []string) error {
	_ = "STUB: not implemented"
	// Render the result values, using HTMX OOB
	return nil
}

func (a *Action) renderResultsDownload(w http.ResponseWriter, valuesMap []map[string]any) error {
	_ = "STUB: not implemented"
	// Render the result values, using HTMX OOB
	return nil
}

func (a *Action) renderResultsImage(w http.ResponseWriter, valuesMap []map[string]any) error {
	_ = "STUB: not implemented"
	// Render the result values, using HTMX OOB
	return nil
}

func (a *Action) renderResultsTable(w http.ResponseWriter, valuesMap []map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// Missing value

func (a *Action) renderResultsJson(w http.ResponseWriter, valuesMap []map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func RunDeferredCleanup(thread *starlark.Thread) error { _ = "STUB: not implemented"; return nil }

// reset the defer map

type ParamDef struct {
	Name               string
	Description        string
	Value              any
	InputType          string
	Options            []string
	DisplayType        string
	DisplayTypeOptions string
}

const (
	OPTIONS_PREFIX            = "options-"
	OPTIONS_PREFIX_UNDERSCORE = "options_"
	// Both are allowed, for backward compatibility. Underscore is preferred since it is a valid starlark identifier
)

func (a *Action) getForm(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// params with options-x prefix are treated as select options for x

// Prefer value from query params

// Default to string format

func (a *Action) getLinksWithQS(ctx context.Context, qs string) []ActionLink {
	_ = "STUB: not implemented"
	return nil
}

// Don't add self link

// whether this user has access to this action

func (a *Action) handleSuggestResponse(ctx context.Context, w http.ResponseWriter, paramQS string, retVal starlark.Value) {
	_ = "STUB: not implemented"
	return
}

// No suggestions available
