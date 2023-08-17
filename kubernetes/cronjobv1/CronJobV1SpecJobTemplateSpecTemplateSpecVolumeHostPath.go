package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecVolumeHostPath struct {
	// Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#path CronJobV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Type for HostPath volume. Allowed values are "" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#type CronJobV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

