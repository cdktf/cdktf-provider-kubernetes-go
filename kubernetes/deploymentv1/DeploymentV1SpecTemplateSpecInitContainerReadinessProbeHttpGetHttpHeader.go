package deploymentv1


type DeploymentV1SpecTemplateSpecInitContainerReadinessProbeHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/deployment_v1#name DeploymentV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/deployment_v1#value DeploymentV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

