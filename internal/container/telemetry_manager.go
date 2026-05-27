// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
)

func WrapContainerManager(cm ContainerManager, kind string) ContainerManager {
	_ = "STUB: not implemented"
	return *new(ContainerManager)
}

type telemetryContainerManager struct {
	ContainerManager
	kind string
}

func (m *telemetryContainerManager) BuildImage(ctx context.Context, name ImageName, sourceUrl, containerFile string, containerArgs map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *telemetryContainerManager) ImageExists(ctx context.Context, name ImageName) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *telemetryContainerManager) RefreshImage(ctx context.Context, name ImageName) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *telemetryContainerManager) GetContainerState(ctx context.Context, name ContainerName, expectHash string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func (m *telemetryContainerManager) StartContainer(ctx context.Context, name ContainerName) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *telemetryContainerManager) StopContainer(ctx context.Context, name ContainerName) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *telemetryContainerManager) RunContainer(ctx context.Context, appEntry *types.AppEntry, sourceDir string, containerName ContainerName,
	imageName ImageName, port int32, envMap map[string]string, volumes []*VolumeInfo,
	containerOptions map[string]string, paramMap map[string]string, versionHash string, isImageSpec bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *telemetryContainerManager) GetContainerLogs(ctx context.Context, name ContainerName, linesToShow int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *telemetryContainerManager) VolumeExists(ctx context.Context, name VolumeName) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *telemetryContainerManager) VolumeCreate(ctx context.Context, name VolumeName) error {
	_ = "STUB: not implemented"
	return nil
}

type telemetryDevContainerManager struct {
	*telemetryContainerManager
	dev DevContainerManager
}

func (m *telemetryDevContainerManager) RemoveImage(ctx context.Context, name ImageName) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *telemetryDevContainerManager) RemoveContainer(ctx context.Context, name ContainerName) error {
	_ = "STUB: not implemented"
	return nil
}
