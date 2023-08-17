package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEphemeral struct {
	// volume_claim_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#volume_claim_template CronJobV1#volume_claim_template}
	VolumeClaimTemplate *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate `field:"required" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

