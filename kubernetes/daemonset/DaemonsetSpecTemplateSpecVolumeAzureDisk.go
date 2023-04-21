package daemonset


type DaemonsetSpecTemplateSpecVolumeAzureDisk struct {
	// Host Caching mode: None, Read Only, Read Write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#caching_mode Daemonset#caching_mode}
	CachingMode *string `field:"required" json:"cachingMode" yaml:"cachingMode"`
	// The URI the data disk in the blob storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#data_disk_uri Daemonset#data_disk_uri}
	DataDiskUri *string `field:"required" json:"dataDiskUri" yaml:"dataDiskUri"`
	// The Name of the data disk in the blob storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#disk_name Daemonset#disk_name}
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#fs_type Daemonset#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The type for the data disk. Expected values: Shared, Dedicated, Managed. Defaults to Shared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#kind Daemonset#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#read_only Daemonset#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

