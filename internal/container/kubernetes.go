// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"context"

	"github.com/openrundev/openrun/internal/types"
	core "k8s.io/api/core/v1"

	corev1apply "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	OPENRUN_FIELD_MANAGER = "openrun"
)

type KubernetesOptions struct {
	Cpus        string         `mapstructure:"cpus"`
	Memory      string         `mapstructure:"memory"`
	MinReplicas int32          `mapstructure:"min_replicas"` // min number of replicas to run the app on
	MaxReplicas int32          `mapstructure:"max_replicas"` // max number of replicas to run the app on
	Other       map[string]any `mapstructure:",remain"`
}

func parseKubernetesOptions(options map[string]string) (KubernetesOptions, error) {
	_ = "STUB: not implemented"
	return *new(KubernetesOptions), nil
}

type KubernetesCM struct {
	*types.Logger
	appNamespace string
	config       *types.ServerConfig
	clientSet    kubernetes.Interface
	restConfig   *rest.Config
	appConfig    *types.AppConfig
	appRunDir    string
	appId        types.AppId
}

func sanitizeContainerName(name string) string { _ = "STUB: not implemented"; return "" }

// max length for a Kubernetes object name is 63, leave space for the suffix

// TrimLabelValue trims the input string to 63 characters so that it can be used as a Kubernetes label value
func TrimLabelValue(input string) string { _ = "STUB: not implemented"; return "" }

// isPodReady checks if a pod is both running and has the Ready condition set to True.
// This is important because Kubernetes Services only route traffic to Ready pods.
func isPodReady(pod *core.Pod) bool { _ = "STUB: not implemented"; return false }

func currentNamespace() (string, error) { _ = "STUB: not implemented"; return "", nil }

func namespaceExists(ctx context.Context, client kubernetes.Interface, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// real error (RBAC, network, etc.)

func NewKubernetesCM(logger *types.Logger, config *types.ServerConfig, appConfig *types.AppConfig, appRunDir string, appId types.AppId) (*KubernetesCM, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ ContainerManager = (*KubernetesCM)(nil)

func loadConfig() (*rest.Config, error) {
	_ = "STUB: not implemented"
	// Try in-cluster; fall back to default kubeconfig
	return nil, nil
}

func (k *KubernetesCM) SupportsInPlaceUpdate() bool { _ = "STUB: not implemented"; return false }

func (k *KubernetesCM) ImageExists(ctx context.Context, name ImageName) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// RefreshImage is a no-op on Kubernetes. Returning an empty digest leaves the existing identity hash and
// container lifecycle on Kubernetes unchanged.
func (k *KubernetesCM) RefreshImage(ctx context.Context, name ImageName) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *KubernetesCM) BuildImage(ctx context.Context, imgName ImageName, sourceUrl, containerFile string, containerArgs map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate Docker config JSON only for Kaniko (which needs it as a Kubernetes secret)

func (k *KubernetesCM) GetContainerState(ctx context.Context, name ContainerName, expectHash string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// version hash mismatch, deployment is not in the correct state

// Get the pods which are part of this deployment

func (k *KubernetesCM) StartContainer(ctx context.Context, name ContainerName) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KubernetesCM) StopContainer(ctx context.Context, name ContainerName) error {
	_ = "STUB: not implemented"
	return nil
}

// scale down to zero

func (k *KubernetesCM) RunContainer(ctx context.Context, appEntry *types.AppEntry, sourceDir string, containerName ContainerName,
	imageName ImageName, port int32, envMap map[string]string, volumes []*VolumeInfo,
	containerOptions map[string]string, paramMap map[string]string, versionHash string, isImageSpec bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KubernetesCM) GetContainerLogs(ctx context.Context, name ContainerName, linesToShow int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List pods with the matching label

// Get the first pod

// Get logs from the first container

//nolint:errcheck

func (k *KubernetesCM) VolumeExists(ctx context.Context, name VolumeName) bool {
	_ = "STUB: not implemented"
	return false
}

func (k *KubernetesCM) VolumeCreate(ctx context.Context, name VolumeName) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: support other access modes

// processVolumes converts VolumeInfo entries to Kubernetes Volume and VolumeMount configurations.
// It creates Secrets for secret volumes, ConfigMaps for volumes without a VolumeName, and
// references existing PVCs for named volumes.
func (k *KubernetesCM) processVolumes(ctx context.Context, name string, volumes []*VolumeInfo, sourceDir string, paramMap map[string]string) (
	[]*corev1apply.VolumeApplyConfiguration, []*corev1apply.VolumeMountApplyConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create a Secret from the source file and mount it

// Create a ConfigMap from the source file and mount it

// unnamed volume, use the path for generating the volume name

// PVC-based volume (already created via VolumeCreate)

// use the same name for the volume reference

const VERSION_HASH_LABEL = LABEL_PREFIX + "version.hash"

// createDeployment creates a Deployment + Service using server-side apply and returns the Service URL.
func (k *KubernetesCM) createDeployment(ctx context.Context, name, image string,
	port int32, envMap map[string]string, volumes []*VolumeInfo, sourceDir string, paramMap map[string]string,
	appEntry *types.AppEntry, versionHash string, kubernetesOptions KubernetesOptions, isImageSpec bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// For image-spec apps, the version hash does not change when only the
// upstream tag's content moves (Kubernetes's RefreshImage is a no-op so
// the image digest is not folded into versionHash). Without something
// changing in the pod template, a re-apply during `openrun app reload`
// would be a no-op and Kubernetes would not roll out new pods. We bump
// a dedicated pod-template-only annotation on every reload so the apply
// diff exists and the RollingUpdate strategy picks up the new image. We
// keep this annotation off the Deployment- and Service-level metadata
// so unrelated controllers don't see spurious churn.

// Set replicas from kubernetesOptions, defaulting to 1

// Convert envMap to Kubernetes EnvVar apply configurations

// Process volumes (creates Secrets/ConfigMaps as needed)

// Image-spec apps consume an externally-managed tag; ensure each new
// pod actually re-resolves the tag against the registry instead of
// reusing whatever bytes happen to be cached on the node. Combined
// with the refreshed-at annotation bump above, every reload pulls
// the latest content for the configured tag.

// Add resource requirements if cpus or memory are specified

// convert to millicores

// Set deployment strategy

//WithType(appsv1.RecreateDeploymentStrategyType)

// Create HPA if MaxReplicas > 1

// In-cluster DNS URL
