// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumev1


type PersistentVolumeV1SpecPersistentVolumeSourceAzureDisk struct {
	// Host Caching mode: None, Read Only, Read Write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#caching_mode PersistentVolumeV1#caching_mode}
	CachingMode *string `field:"required" json:"cachingMode" yaml:"cachingMode"`
	// The URI the data disk in the blob storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#data_disk_uri PersistentVolumeV1#data_disk_uri}
	DataDiskUri *string `field:"required" json:"dataDiskUri" yaml:"dataDiskUri"`
	// The Name of the data disk in the blob storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#disk_name PersistentVolumeV1#disk_name}
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#fs_type PersistentVolumeV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The type for the data disk. Expected values: Shared, Dedicated, Managed. Defaults to Shared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#kind PersistentVolumeV1#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#read_only PersistentVolumeV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

