package daemonset


type DaemonsetSpecStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemonset#rolling_update Daemonset#rolling_update}
	RollingUpdate *DaemonsetSpecStrategyRollingUpdate `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemonset#type Daemonset#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

