package statefulsetv1


type StatefulSetV1SpecTemplateSpecVolumeEmptyDir struct {
	// What type of storage medium should back this directory.
	//
	// The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set_v1#medium StatefulSetV1#medium}
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// Total amount of local storage required for this EmptyDir volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set_v1#size_limit StatefulSetV1#size_limit}
	SizeLimit *string `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

