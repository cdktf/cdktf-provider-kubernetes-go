package daemonsetv1


type DaemonSetV1SpecTemplateSpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#name DaemonSetV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#value DaemonSetV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

