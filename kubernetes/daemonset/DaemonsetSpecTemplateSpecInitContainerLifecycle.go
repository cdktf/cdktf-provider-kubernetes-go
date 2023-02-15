package daemonset


type DaemonsetSpecTemplateSpecInitContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#post_start Daemonset#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#pre_stop Daemonset#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

