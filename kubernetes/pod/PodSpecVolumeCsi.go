// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecVolumeCsi struct {
	// the name of the volume driver to use. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod#driver Pod#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod#fs_type Pod#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// node_publish_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod#node_publish_secret_ref Pod#node_publish_secret_ref}
	NodePublishSecretRef *PodSpecVolumeCsiNodePublishSecretRef `field:"optional" json:"nodePublishSecretRef" yaml:"nodePublishSecretRef"`
	// Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". More info: https://kubernetes.io/docs/concepts/storage/volumes#csi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod#read_only Pod#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Attributes of the volume to publish.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod#volume_attributes Pod#volume_attributes}
	VolumeAttributes *map[string]*string `field:"optional" json:"volumeAttributes" yaml:"volumeAttributes"`
}

