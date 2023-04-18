package daemonsetv1


type DaemonSetV1SpecStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemon_set_v1#rolling_update DaemonSetV1#rolling_update}
	RollingUpdate *DaemonSetV1SpecStrategyRollingUpdate `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemon_set_v1#type DaemonSetV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

