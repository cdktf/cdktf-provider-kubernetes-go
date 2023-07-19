package cronjob


type CronJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#delete CronJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

