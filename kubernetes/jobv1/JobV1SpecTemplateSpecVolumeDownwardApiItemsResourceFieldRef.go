package jobv1


type JobV1SpecTemplateSpecVolumeDownwardApiItemsResourceFieldRef struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#container_name JobV1#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#resource JobV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#divisor JobV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

