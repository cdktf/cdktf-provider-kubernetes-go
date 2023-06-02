package deployment


type DeploymentSpecTemplateSpecVolumePhotonPersistentDisk struct {
	// ID that identifies Photon Controller persistent disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/deployment#pd_id Deployment#pd_id}
	PdId *string `field:"required" json:"pdId" yaml:"pdId"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/deployment#fs_type Deployment#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

