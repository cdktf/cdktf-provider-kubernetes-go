package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#name CronJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#value CronJob#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

