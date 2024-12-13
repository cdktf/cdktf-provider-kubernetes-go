// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecTemplateSpecVolumeFlexVolume struct {
	// Driver is the name of the driver to use for this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/stateful_set_v1#driver StatefulSetV1#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/stateful_set_v1#fs_type StatefulSetV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Extra command options if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/stateful_set_v1#options StatefulSetV1#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/stateful_set_v1#read_only StatefulSetV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/stateful_set_v1#secret_ref StatefulSetV1#secret_ref}
	SecretRef *StatefulSetV1SpecTemplateSpecVolumeFlexVolumeSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

