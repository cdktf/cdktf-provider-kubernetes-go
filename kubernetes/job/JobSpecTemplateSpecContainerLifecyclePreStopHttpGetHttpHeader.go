package job


type JobSpecTemplateSpecContainerLifecyclePreStopHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#name Job#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#value Job#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

