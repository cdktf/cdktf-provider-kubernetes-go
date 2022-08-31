// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type CronJobSpecJobTemplateSpecTemplateSpecInitContainerLifecyclePostStartHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#name CronJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#value CronJob#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

