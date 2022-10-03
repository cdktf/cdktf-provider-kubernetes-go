package cronjobv1


type CronJobV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#delete CronJobV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

