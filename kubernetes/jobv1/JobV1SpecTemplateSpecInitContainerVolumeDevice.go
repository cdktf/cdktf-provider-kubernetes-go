// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecTemplateSpecInitContainerVolumeDevice struct {
	// Path within the container at which the volume device should be attached. For example '/dev/xvda'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#device_path JobV1#device_path}
	DevicePath *string `field:"required" json:"devicePath" yaml:"devicePath"`
	// This must match the Name of a PersistentVolumeClaim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#name JobV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

