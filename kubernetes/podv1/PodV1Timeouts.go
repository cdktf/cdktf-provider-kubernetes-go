package podv1


type PodV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#create PodV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#delete PodV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

