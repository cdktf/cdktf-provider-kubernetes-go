package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeral struct {
	// volume_claim_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job#volume_claim_template CronJob#volume_claim_template}
	VolumeClaimTemplate *CronJobSpecJobTemplateSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate `field:"required" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

