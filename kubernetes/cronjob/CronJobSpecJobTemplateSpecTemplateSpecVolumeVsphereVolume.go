package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolumeVsphereVolume struct {
	// Path that identifies vSphere volume vmdk.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#volume_path CronJob#volume_path}
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#fs_type CronJob#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
}

