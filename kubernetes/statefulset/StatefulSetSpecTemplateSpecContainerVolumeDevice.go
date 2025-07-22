// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecContainerVolumeDevice struct {
	// Path within the container at which the volume device should be attached. For example '/dev/xvda'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set#device_path StatefulSet#device_path}
	DevicePath *string `field:"required" json:"devicePath" yaml:"devicePath"`
	// This must match the Name of a PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set#name StatefulSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

