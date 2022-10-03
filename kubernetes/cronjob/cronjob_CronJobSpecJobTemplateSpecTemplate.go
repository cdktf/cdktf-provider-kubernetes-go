package cronjob


type CronJobSpecJobTemplateSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#metadata CronJob#metadata}
	Metadata *CronJobSpecJobTemplateSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#spec CronJob#spec}
	Spec *CronJobSpecJobTemplateSpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

