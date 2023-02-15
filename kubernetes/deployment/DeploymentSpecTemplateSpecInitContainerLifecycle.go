package deployment


type DeploymentSpecTemplateSpecInitContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#post_start Deployment#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#pre_stop Deployment#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

