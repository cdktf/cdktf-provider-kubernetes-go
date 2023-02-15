package persistentvolumev1


type PersistentVolumeV1SpecPersistentVolumeSourceCsi struct {
	// the name of the volume driver to use. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#driver PersistentVolumeV1#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// A string value that uniquely identifies the volume. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#volume_handle PersistentVolumeV1#volume_handle}
	VolumeHandle *string `field:"required" json:"volumeHandle" yaml:"volumeHandle"`
	// controller_expand_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#controller_expand_secret_ref PersistentVolumeV1#controller_expand_secret_ref}
	ControllerExpandSecretRef *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerExpandSecretRef `field:"optional" json:"controllerExpandSecretRef" yaml:"controllerExpandSecretRef"`
	// controller_publish_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#controller_publish_secret_ref PersistentVolumeV1#controller_publish_secret_ref}
	ControllerPublishSecretRef *PersistentVolumeV1SpecPersistentVolumeSourceCsiControllerPublishSecretRef `field:"optional" json:"controllerPublishSecretRef" yaml:"controllerPublishSecretRef"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#fs_type PersistentVolumeV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// node_publish_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#node_publish_secret_ref PersistentVolumeV1#node_publish_secret_ref}
	NodePublishSecretRef *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodePublishSecretRef `field:"optional" json:"nodePublishSecretRef" yaml:"nodePublishSecretRef"`
	// node_stage_secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#node_stage_secret_ref PersistentVolumeV1#node_stage_secret_ref}
	NodeStageSecretRef *PersistentVolumeV1SpecPersistentVolumeSourceCsiNodeStageSecretRef `field:"optional" json:"nodeStageSecretRef" yaml:"nodeStageSecretRef"`
	// Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". More info: http://kubernetes.io/docs/user-guide/volumes#csi.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#read_only PersistentVolumeV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Attributes of the volume to publish.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#volume_attributes PersistentVolumeV1#volume_attributes}
	VolumeAttributes *map[string]*string `field:"optional" json:"volumeAttributes" yaml:"volumeAttributes"`
}

