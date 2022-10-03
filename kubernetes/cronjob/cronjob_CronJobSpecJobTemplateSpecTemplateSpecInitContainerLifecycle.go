package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecInitContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#post_start CronJob#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#pre_stop CronJob#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

