// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type CronJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#delete CronJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

