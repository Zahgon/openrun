// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"context"
	"encoding/base32"
	"sync"

	"github.com/openrundev/openrun/internal/types"
)

var base32encoder = base32.StdEncoding.WithPadding(base32.NoPadding)

func genLowerCaseId(name string) string {
	_ = "STUB: not implemented"
	// The container id needs to be lower case. Use base32 to encode the name so that it can be lowercased
	return ""
}

var mu sync.Mutex
var buildLockChannel chan string // channel to hold the build ids, max size is MaxConcurrentBuilds

// acquireBuildLock acquires a build lock for the given build id. If the lock is not available,
// it will wait for the lock to be available or the context to be done.
// The lock is released when the returned function is called.
func acquireBuildLock(ctx context.Context, config *types.SystemConfig, buildId string) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CommandOptions struct {
	Cpus   string         `mapstructure:"cpus"`
	Memory string         `mapstructure:"memory"`
	Other  map[string]any `mapstructure:",remain"`
}

func parseCommandOptions(command string, options map[string]string) (CommandOptions, error) {
	_ = "STUB: not implemented"
	return *new(CommandOptions), nil
}

func ParseCommandOptions(containerCommand string, options map[string]string) (CommandOptions, error) {
	_ = "STUB: not implemented"
	return *new(CommandOptions), nil
}

type CommandCM struct {
	*types.Logger
	appRunDir string
	appId     types.AppId
	config    *types.ServerConfig
}

var _ DevContainerManager = (*CommandCM)(nil)

func NewCommandCM(logger *types.Logger, config *types.ServerConfig, appId types.AppId, appRunDir string) *CommandCM {
	_ = "STUB: not implemented"
	return nil
}

func (k *CommandCM) SupportsInPlaceUpdate() bool { _ = "STUB: not implemented"; return false }

func (c *CommandCM) RemoveImage(ctx context.Context, name ImageName) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommandCM) BuildImage(ctx context.Context, imgName ImageName, sourceUrl, containerFile string, containerArgs map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildImageCommand(ctx context.Context, logger *types.Logger, config *types.ServerConfig,
	imgName ImageName, sourceUrl, containerFile string, containerArgs map[string]string, containerCommand string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommandCM) RemoveContainer(ctx context.Context, name ContainerName) error {
	_ = "STUB: not implemented"
	return nil
}

// GetContainerState returns the host:port of the running container, "" if not running. running is true if the container is running.
func (c *CommandCM) GetContainerState(ctx context.Context, name ContainerName, expectHash string) (string, bool, error) {
	_ = "STUB: not implemented"
	// expectedHash is ignored for command based container manager, since it does not do in place updates
	return "", false, nil
}

// version hash is not used for command based container manager, since it does not do in place updates

func (c *CommandCM) getContainers(ctx context.Context, name ContainerName, getAll bool) ([]Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListOpenRunContainers returns running containers with an OpenRun ownership label.
func (c *CommandCM) ListOpenRunContainers(ctx context.Context) ([]Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listContainers runs `<containerCommand> ps --format json` with the given
// filters and parses the result. Handles both Podman (JSON array, Names/Ports
// as arrays) and Docker (newline-separated JSON objects).
func (c *CommandCM) listContainers(ctx context.Context, filters []string, getAll bool) ([]Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck
// Podman format (Names and Ports are arrays)

// only HostPort is needed

// JSON output (podman)

// Newline separated JSON (Docker)

// "Ports":"127.0.0.1:55000-\u003e5000/tcp"

func (c *CommandCM) GetContainerLogs(ctx context.Context, name ContainerName, linesToShow int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *CommandCM) StopContainer(ctx context.Context, name ContainerName) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommandCM) StartContainer(ctx context.Context, name ContainerName) error {
	_ = "STUB: not implemented"
	return nil
}

const LABEL_PREFIX = "dev.openrun."

func (c *CommandCM) RunContainer(ctx context.Context, appEntry *types.AppEntry, sourceDir string, containerName ContainerName,
	imageName ImageName, port int32, envMap map[string]string, volumes []*VolumeInfo,
	containerOptions map[string]string, paramMap map[string]string, versionHash string, isImageSpec bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Add env args

// Add container related args

// RefreshImage pulls the named image and returns its content-addressable
// digest. It first attempts to extract the manifest digest from RepoDigests
// (which is stable across container managers and matches the digest the
// registry advertises); it falls back to the image config digest (.Id) when
// the local image has no associated RepoDigests entry (e.g. it was built
// locally rather than pulled).
func (c *CommandCM) RefreshImage(ctx context.Context, name ImageName) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RepoDigests entries are "repo/name@sha256:abc..."; strip the repo prefix.

func (c *CommandCM) ImageExists(ctx context.Context, name ImageName) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ExecTailN executes a command and returns the last n lines of output
func (c *CommandCM) ExecTailN(ctx context.Context, command string, args []string, n int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a ring buffer to hold the last 1000 lines of output

// Push the latest line into the ring buffer, displacing the oldest line if necessary

func (c CommandCM) VolumeExists(ctx context.Context, name VolumeName) bool {
	_ = "STUB: not implemented"
	return false
}

func (c CommandCM) VolumeCreate(ctx context.Context, name VolumeName) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommandCM) genMountArgs(sourceDir string, volumeInfo []*VolumeInfo, paramMap map[string]string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For cl_secret:file.prop:/data/file.prop, pass file.prop through the template
// processor, write output to file.prop.gen and then bind mount it as
// /source_dir/file.prop.gen:/data/file.prop

// bind mount

// unnamed volume, use the path for generating the volume name

const (
	DOCKER_COMMAND = "docker"
	PODMAN_COMMAND = "podman"

	kubeHostEnv = "KUBERNETES_SERVICE_HOST"
	kubePortEnv = "KUBERNETES_SERVICE_PORT"
)

func LookupContainerCommand(checkKubernetes bool) string { _ = "STUB: not implemented"; return "" }

// Check if running in Kubernetes
