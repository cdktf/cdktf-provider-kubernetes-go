// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type CronJobV1SpecJobTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#metadata CronJobV1#metadata}
	Metadata *CronJobV1SpecJobTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#spec CronJobV1#spec}
	Spec *CronJobV1SpecJobTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}
