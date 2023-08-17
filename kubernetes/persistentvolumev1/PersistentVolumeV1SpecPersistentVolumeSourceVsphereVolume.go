package persistentvolumev1


type PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume struct {
	// Path that identifies vSphere volume vmdk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume_v1#volume_path PersistentVolumeV1#volume_path}
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume_v1#fs_type PersistentVolumeV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

