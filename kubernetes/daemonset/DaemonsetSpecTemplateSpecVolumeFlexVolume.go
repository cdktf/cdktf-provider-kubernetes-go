package daemonset


type DaemonsetSpecTemplateSpecVolumeFlexVolume struct {
	// Driver is the name of the driver to use for this volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#driver Daemonset#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#fs_type Daemonset#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Extra command options if any.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#options Daemonset#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#read_only Daemonset#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#secret_ref Daemonset#secret_ref}
	SecretRef *DaemonsetSpecTemplateSpecVolumeFlexVolumeSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

