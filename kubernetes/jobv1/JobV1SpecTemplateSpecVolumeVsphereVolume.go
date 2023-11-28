// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecTemplateSpecVolumeVsphereVolume struct {
	// Path that identifies vSphere volume vmdk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/job_v1#volume_path JobV1#volume_path}
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/job_v1#fs_type JobV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

