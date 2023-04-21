package podv1


type PodV1SpecInitContainerEnv struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#name PodV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables.
	//
	// If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#value PodV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#value_from PodV1#value_from}
	ValueFrom *PodV1SpecInitContainerEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

