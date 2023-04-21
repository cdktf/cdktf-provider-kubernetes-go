package deployment


type DeploymentSpecTemplateSpecInitContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/deployment#post_start Deployment#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/deployment#pre_stop Deployment#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

