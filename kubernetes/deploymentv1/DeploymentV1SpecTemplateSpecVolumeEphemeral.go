package deploymentv1


type DeploymentV1SpecTemplateSpecVolumeEphemeral struct {
	// volume_claim_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment_v1#volume_claim_template DeploymentV1#volume_claim_template}
	VolumeClaimTemplate *DeploymentV1SpecTemplateSpecVolumeEphemeralVolumeClaimTemplate `field:"required" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

