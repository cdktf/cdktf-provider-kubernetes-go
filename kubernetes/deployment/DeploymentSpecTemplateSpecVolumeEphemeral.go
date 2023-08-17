package deployment


type DeploymentSpecTemplateSpecVolumeEphemeral struct {
	// volume_claim_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment#volume_claim_template Deployment#volume_claim_template}
	VolumeClaimTemplate *DeploymentSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate `field:"required" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

