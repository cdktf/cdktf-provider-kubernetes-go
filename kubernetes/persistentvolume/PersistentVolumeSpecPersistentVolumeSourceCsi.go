// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolume


type PersistentVolumeSpecPersistentVolumeSourceCsi struct {
	// the name of the volume driver to use. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#driver PersistentVolume#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// A string value that uniquely identifies the volume. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#volume_handle PersistentVolume#volume_handle}
	VolumeHandle *string `field:"required" json:"volumeHandle" yaml:"volumeHandle"`
	// controller_expand_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#controller_expand_secret_ref PersistentVolume#controller_expand_secret_ref}
	ControllerExpandSecretRef *PersistentVolumeSpecPersistentVolumeSourceCsiControllerExpandSecretRef `field:"optional" json:"controllerExpandSecretRef" yaml:"controllerExpandSecretRef"`
	// controller_publish_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#controller_publish_secret_ref PersistentVolume#controller_publish_secret_ref}
	ControllerPublishSecretRef *PersistentVolumeSpecPersistentVolumeSourceCsiControllerPublishSecretRef `field:"optional" json:"controllerPublishSecretRef" yaml:"controllerPublishSecretRef"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#fs_type PersistentVolume#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// node_publish_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#node_publish_secret_ref PersistentVolume#node_publish_secret_ref}
	NodePublishSecretRef *PersistentVolumeSpecPersistentVolumeSourceCsiNodePublishSecretRef `field:"optional" json:"nodePublishSecretRef" yaml:"nodePublishSecretRef"`
	// node_stage_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#node_stage_secret_ref PersistentVolume#node_stage_secret_ref}
	NodeStageSecretRef *PersistentVolumeSpecPersistentVolumeSourceCsiNodeStageSecretRef `field:"optional" json:"nodeStageSecretRef" yaml:"nodeStageSecretRef"`
	// Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". More info: https://kubernetes.io/docs/concepts/storage/volumes#csi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#read_only PersistentVolume#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Attributes of the volume to publish.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/persistent_volume#volume_attributes PersistentVolume#volume_attributes}
	VolumeAttributes *map[string]*string `field:"optional" json:"volumeAttributes" yaml:"volumeAttributes"`
}

