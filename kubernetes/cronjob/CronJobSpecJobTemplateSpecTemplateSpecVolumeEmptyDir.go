package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolumeEmptyDir struct {
	// What type of storage medium should back this directory.
	//
	// The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job#medium CronJob#medium}
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// Total amount of local storage required for this EmptyDir volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job#size_limit CronJob#size_limit}
	SizeLimit *string `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

