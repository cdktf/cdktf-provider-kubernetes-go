package pod


type PodSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#container_name Pod#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#resource Pod#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#divisor Pod#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

