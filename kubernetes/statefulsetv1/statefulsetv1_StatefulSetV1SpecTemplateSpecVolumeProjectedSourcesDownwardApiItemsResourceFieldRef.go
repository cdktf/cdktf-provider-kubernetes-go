package statefulsetv1


type StatefulSetV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#container_name StatefulSetV1#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#resource StatefulSetV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#divisor StatefulSetV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}
