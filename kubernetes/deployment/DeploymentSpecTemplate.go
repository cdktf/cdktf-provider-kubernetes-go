package deployment


type DeploymentSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/deployment#metadata Deployment#metadata}
	Metadata *DeploymentSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/deployment#spec Deployment#spec}
	Spec *DeploymentSpecTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}
