// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// preference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#preference CronJobV1#preference}
	Preference *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	// weight is in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#weight CronJobV1#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}
