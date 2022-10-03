package deployment


type DeploymentSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#metadata Deployment#metadata}
	Metadata *DeploymentSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#spec Deployment#spec}
	Spec *DeploymentSpecTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

