// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller


type ReplicationControllerSpecTemplateSpecVolumeFlexVolume struct {
	// Driver is the name of the driver to use for this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller#driver ReplicationController#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller#fs_type ReplicationController#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Extra command options if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller#options ReplicationController#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller#read_only ReplicationController#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller#secret_ref ReplicationController#secret_ref}
	SecretRef *ReplicationControllerSpecTemplateSpecVolumeFlexVolumeSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

