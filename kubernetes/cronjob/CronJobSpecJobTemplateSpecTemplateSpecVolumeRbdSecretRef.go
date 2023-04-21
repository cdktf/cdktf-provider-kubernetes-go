package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolumeRbdSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job#name CronJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job#namespace CronJob#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

