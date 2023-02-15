package pod


type PodSpecVolumePhotonPersistentDisk struct {
	// ID that identifies Photon Controller persistent disk.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#pd_id Pod#pd_id}
	PdId *string `field:"required" json:"pdId" yaml:"pdId"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#fs_type Pod#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

