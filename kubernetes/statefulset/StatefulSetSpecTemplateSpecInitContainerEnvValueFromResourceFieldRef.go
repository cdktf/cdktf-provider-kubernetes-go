package statefulset


type StatefulSetSpecTemplateSpecInitContainerEnvValueFromResourceFieldRef struct {
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#resource StatefulSet#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#container_name StatefulSet#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#divisor StatefulSet#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

