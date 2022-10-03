package cronjob


type CronJobSpecJobTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#metadata CronJob#metadata}
	Metadata *CronJobSpecJobTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#spec CronJob#spec}
	Spec *CronJobSpecJobTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

