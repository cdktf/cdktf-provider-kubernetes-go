package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecInitContainerEnvFromSecretRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#name CronJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the Secret must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#optional CronJob#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

