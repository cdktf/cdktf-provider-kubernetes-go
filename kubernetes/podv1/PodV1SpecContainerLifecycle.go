package podv1


type PodV1SpecContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#post_start PodV1#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#pre_stop PodV1#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

