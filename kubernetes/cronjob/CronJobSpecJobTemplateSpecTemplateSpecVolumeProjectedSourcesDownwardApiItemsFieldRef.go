package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef struct {
	// Version of the schema the FieldPath is written in terms of, defaults to 'v1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/cron_job#api_version CronJob#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Path of the field to select in the specified API version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/cron_job#field_path CronJob#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
}
