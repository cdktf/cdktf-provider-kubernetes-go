// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecContainerVolumeDevice struct {
	// Path within the container at which the volume device should be attached. For example '/dev/xvda'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod#device_path Pod#device_path}
	DevicePath *string `field:"required" json:"devicePath" yaml:"devicePath"`
	// This must match the Name of a PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod#name Pod#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

