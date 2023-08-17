package persistentvolume


type PersistentVolumeSpecPersistentVolumeSourceCinder struct {
	// Volume ID used to identify the volume in Cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#volume_id PersistentVolume#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#fs_type PersistentVolume#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). More info: https://examples.k8s.io/mysql-cinder-pd/README.md.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume#read_only PersistentVolume#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

