package podv1


type PodV1SpecInitContainerLifecyclePreStopHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod_v1#name PodV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod_v1#value PodV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

