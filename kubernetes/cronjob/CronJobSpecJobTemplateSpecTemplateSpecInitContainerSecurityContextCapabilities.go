package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecInitContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job#add CronJob#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job#drop CronJob#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

