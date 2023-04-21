package statefulsetv1


type StatefulSetV1SpecTemplateSpecInitContainerLifecyclePreStopHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#name StatefulSetV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#value StatefulSetV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

