package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjected struct {
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#sources CronJobV1#sources}
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// Optional: mode bits to use on created files by default.
	//
	// Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#default_mode CronJobV1#default_mode}
	DefaultMode *string `field:"optional" json:"defaultMode" yaml:"defaultMode"`
}

