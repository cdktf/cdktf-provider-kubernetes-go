package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#add CronJobV1#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#drop CronJobV1#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

