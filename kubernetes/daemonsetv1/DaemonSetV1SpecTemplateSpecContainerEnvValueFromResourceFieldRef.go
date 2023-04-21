package daemonsetv1


type DaemonSetV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef struct {
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemon_set_v1#resource DaemonSetV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemon_set_v1#container_name DaemonSetV1#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemon_set_v1#divisor DaemonSetV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

