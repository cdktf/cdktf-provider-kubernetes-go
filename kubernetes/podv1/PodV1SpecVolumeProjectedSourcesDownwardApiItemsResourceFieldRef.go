package podv1


type PodV1SpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#container_name PodV1#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#resource PodV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#divisor PodV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

