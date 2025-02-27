// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolume


type PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDisk struct {
	// ID that identifies Photon Controller persistent disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/persistent_volume#pd_id PersistentVolume#pd_id}
	PdId *string `field:"required" json:"pdId" yaml:"pdId"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/persistent_volume#fs_type PersistentVolume#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

