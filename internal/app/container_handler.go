// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"os/exec"
	"sync"
	"time"

	"github.com/openrundev/openrun/internal/app/appfs"
	"github.com/openrundev/openrun/internal/container"
	"github.com/openrundev/openrun/internal/types"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

type ContainerState string

const (
	ContainerStateUnknown       ContainerState = "unknown"
	ContainerStateRunning       ContainerState = "running"
	ContainerStateIdleShutdown  ContainerState = "idle_shutdown"
	ContainerStateHealthFailure ContainerState = "health_failure"
)

type ContainerHandler struct {
	*types.Logger
	manager         container.ContainerManager
	app             *App
	serverConfig    *types.ServerConfig
	containerFile   string
	image           string              // image name as specified
	GenImageName    container.ImageName // generated image name
	port            int32               // Port number within the container
	hostNamePort    string              // host name : port number for the container
	lifetime        string
	scheme          string
	health          string
	buildDir        string
	sourceFS        appfs.ReadableFS
	paramMap        map[string]string
	volumeInfo      []*container.VolumeInfo
	containerConfig types.Container
	excludeGlob     []string

	// Idle shutdown related fields
	idleShutdownTicker  *time.Ticker
	stateLock           sync.RWMutex
	currentState        ContainerState
	activeContainerName container.ContainerName
	// imageDigest is the digest (e.g. "sha256:abc...") resolved by the most
	// recent RefreshImage call during ProdReload. Only set for image-spec
	// apps. Folded into getAppHash so a moved upstream tag forces a recreate
	// and used to digest-pin the running container. Guarded by stateLock.
	imageDigest string

	// Health check related fields
	healthCheckTicker *time.Ticker
	stripAppPath      bool
	mountArgs         []string
	cargs             map[string]string
	proxyTracker      *Tracker // Track bytes sent and received by the proxy

	envMap     map[string]string
	envMapHash string
	bindings   []*types.Binding
}

func NewContainerHandler(logger *types.Logger, app *App, containerFile string,
	serverConfig *types.ServerConfig, configPort int32, lifetime, scheme, health, buildDir string, sourceFS appfs.ReadableFS,
	paramMap map[string]string, containerConfig types.Container, stripAppPath bool,
	containerVolumes []string, secretsAllowed [][]string, cargs map[string]any, bindings []*types.Binding) (*ContainerHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Using an image

// Using a container file

// Loop through the parsed result to find the EXPOSE and VOLUME instructions

// Can fail if value is an arg like $PORT

// No port configured in app config, use the one from the container file

// Evaluate secrets in the paramMap

// remove the secrets entry, which is a list of secrets the container is allowed to use

// Evaluate secrets in the build args

// Start the idle shutdown check

// Start the health check goroutine

// a.UsesHtmlTemplate is set in initRouter, so it cannot be used here

const (
	VOL_PREFIX_SECRET = "cl_secret:"
)

func dedupVolumes(volumes []string) []string { _ = "STUB: not implemented"; return nil }

// skip the stripped string, keep only the unstripped version

// already seen, skip

func (h *ContainerHandler) idleAppShutdown(ctx context.Context) { _ = "STUB: not implemented"; return }

// Not idle

// Notify the server to close the app so that it gets reinitialized on next API call

func (h *ContainerHandler) healthChecker(ctx context.Context) { _ = "STUB: not implemented"; return }

// wait for 1 minute to let the app start up

// Notify the server to close the app so that it gets reinitialized on next API call

func extractVolumes(node *parser.Node) []string { _ = "STUB: not implemented"; return nil }

func (h *ContainerHandler) GetProxyUrl() string { _ = "STUB: not implemented"; return "" }

func (h *ContainerHandler) GetHealthUrl(appHealthUrl string) string {
	_ = "STUB: not implemented"
	return ""
}

// Health check URL is specified in the app code, use that

func getMapHash(input map[string]string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Sort the keys to ensure consistent hash

// Default to string

func getSliceHash(input []string) (string, error) {
	_ = "STUB: not implemented"
	// Sort the keys to ensure consistent hash
	return "", nil
}

func (h *ContainerHandler) getEnvMap() map[string]string { _ = "STUB: not implemented"; return nil }

// Add the port number to use into the env
// Using PORT instead of CL_PORT since that seems to be the most common convention across apps

// Add the binding environment variables to the env map

func (h *ContainerHandler) getEnvMapAndHash() (map[string]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (h *ContainerHandler) createSpecFiles() ([]string, error) {
	_ = "STUB: not implemented"
	// Create the spec files if they are not already present
	return nil, nil
}

func (h *ContainerHandler) createVolumes(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// bind mount

// unnamed volume, use the path for generating the volume name

func parseBindPaths(vol string) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (h *ContainerHandler) validateVolumeSource(src string, sourceRelative bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (h *ContainerHandler) parseVolumeString(vol string) (*container.VolumeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Secret passed through bind mount

// Bind mount

// Named volume

// Unnamed volume

func (h *ContainerHandler) DevReload(ctx context.Context, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// The image could be rebuild in case of a dry run, without touching the container.
// But a temp image id will have to be used to avoid conflict with the existing image.
// Dryrun is a no-op for now for containers

// Using a container file, rebuild the image

// Don't remove the spec files, it is good if they are checked into the source repo
// Makes the app independent of changes in the spec files

// Create named volumes for the container

// Command lifetime, service is not started, commands will be run with the image

func (h *ContainerHandler) WaitForHealth(attempts int, containerName container.ContainerName, expectHash string) error {
	_ = "STUB: not implemented"
	return nil
}

// Apps like Streamlit require the app path to be present

//nolint:errcheck

func (h *ContainerHandler) getAppHash() (string, error) { _ = "STUB: not implemented"; return "", nil }

// For image-spec apps where RefreshImage has resolved a digest, fold the
// digest into the identity hash so that a moved tag (e.g.
// mycompany/jp-app:latest pointing to new content) yields a different
// container name and forces ProdReload down the recreate path.
//
// The digest is only appended when non-empty so that:
//   1. non-image apps produce the exact same hash as before this change
//      (no spurious container rebuild on upgrade), and
//   2. image-spec apps whose manager returns no digest (e.g. Kubernetes,
//      where RefreshImage is currently a no-op) also keep their existing
//      hash and container lifecycle.

// IsImageSpec reports whether this container handler was configured with an
// upstream image reference (i.e. `--spec image` / `container.source = "image:..."`).
// Such apps need ProdReload to run on every admin reload so that RefreshImage
// can resolve the current digest and recreate the container when the upstream
// tag has moved; build-spec apps only need ProdReload on Initialize since
// their image identity is captured by the source-content hash.
func (h *ContainerHandler) IsImageSpec() bool { _ = "STUB: not implemented"; return false }

// ActiveContainerName returns the last container this handler successfully started or reused.
func (h *ContainerHandler) ActiveContainerName() (container.ContainerName, bool) {
	_ = "STUB: not implemented"
	return *new(container.ContainerName), false
}

func (h *ContainerHandler) ProdReload(ctx context.Context, dryRun bool) error {
	_ = "STUB: not implemented"
	// For image-spec apps (where the operator supplied an upstream image
	// reference via `--spec image`/`image:`), resolve the current digest
	// from the registry before computing the identity hash. This both
	// ensures the latest content is cached locally (a `docker pull` for the
	// command-based managers) and yields a stable identifier that is folded
	// into fullHash so a moved tag forces a container recreate. Skipped on
	// dry-run since it has external side effects.
	return nil
}

// Digest-pin the image reference we pass to RunContainer/the pod
// template. Subsequent pod restarts or scale-ups will fetch the
// exact digest the operator approved at refresh time, not whatever
// the upstream tag points to at that future moment.

// The image could be rebuild in case of a dry run, without touching the container.
// But a temp image id will have to be used to avoid conflict with the existing image.
// Dryrun is a no-op for now for containers

// For image-spec apps on managers that update workloads in-place
// (i.e. Kubernetes), we deliberately bypass the "service already
// healthy, reuse it" short-circuit and always fall through to
// RunContainer. The upstream tag may have moved while the existing
// Deployment's pod-template hash is unchanged (RefreshImage is a
// no-op on Kubernetes so the digest is not folded into fullHash),
// so re-applying the Deployment is what surfaces the new image:
// createDeployment bumps a pod-template annotation on every reload
// and sets imagePullPolicy=Always, triggering a RollingUpdate that
// pulls the latest image content. Build-spec apps and command-based
// managers (Docker/Podman, where the digest is in fullHash) keep
// the existing reuse fast path.

// Service is present, make sure deployment it is in the correct state

// This does not handle the case where volume list has changed

// TODO handle case where image name is specified and param values change, need to restart container in that case

// Using a container file, build the image if required

// Create named volumes for the container

// Start the container with newly built image

// Command lifetime, service is not started, commands will be run with the image

// Cleanup temp dir after image has been built and mount template file has been generated

func (h *ContainerHandler) Close() error { _ = "STUB: not implemented"; return nil }

func (h *ContainerHandler) Run(ctx context.Context, path string, cmdArgs []string, env []string) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Add env args
}

// Add container related args

func (h *ContainerHandler) getBindingEnv() map[string]string {
	_ = "STUB: not implemented"
	// stage, dev and preview apps use staging binding
	return nil
}

// first binding of postgres type will use POSTGRES_URL, second will use POSTGRES2_URL, etc.
