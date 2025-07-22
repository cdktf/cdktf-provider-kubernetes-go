// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecContainerVolumeDevice struct {
	// Path within the container at which the volume device should be attached. For example '/dev/xvda'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/pod_v1#device_path PodV1#device_path}
	DevicePath *string `field:"required" json:"devicePath" yaml:"devicePath"`
	// This must match the Name of a PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/pod_v1#name PodV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

