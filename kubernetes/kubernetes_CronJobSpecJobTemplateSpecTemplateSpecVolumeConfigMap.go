// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type CronJobSpecJobTemplateSpecTemplateSpecVolumeConfigMap struct {
	// Optional: mode bits to use on created files by default.
	//
	// Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#default_mode CronJob#default_mode}
	DefaultMode *string `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#items CronJob#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#name CronJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional: Specify whether the ConfigMap or its keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#optional CronJob#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}
