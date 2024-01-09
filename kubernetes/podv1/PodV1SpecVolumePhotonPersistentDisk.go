// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecVolumePhotonPersistentDisk struct {
	// ID that identifies Photon Controller persistent disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#pd_id PodV1#pd_id}
	PdId *string `field:"required" json:"pdId" yaml:"pdId"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#fs_type PodV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

