package deployment


type DeploymentSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate struct {
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment#spec Deployment#spec}
	Spec *DeploymentSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment#metadata Deployment#metadata}
	Metadata *DeploymentSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

