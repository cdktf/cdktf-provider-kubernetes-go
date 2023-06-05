package cronjobv1


type CronJobV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/cron_job_v1#delete CronJobV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

