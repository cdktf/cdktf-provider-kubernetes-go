// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type CronJobSpecJobTemplateSpecTemplateSpecContainerEnvValueFromResourceFieldRef struct {
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#resource CronJob#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#container_name CronJob#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#divisor CronJob#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

