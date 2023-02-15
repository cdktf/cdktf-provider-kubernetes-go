package deployment


type DeploymentSpecTemplateSpecContainerReadinessProbeHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#name Deployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#value Deployment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

