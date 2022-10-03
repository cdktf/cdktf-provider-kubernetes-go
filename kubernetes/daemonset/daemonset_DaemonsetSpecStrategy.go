package daemonset


type DaemonsetSpecStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#rolling_update Daemonset#rolling_update}
	RollingUpdate *DaemonsetSpecStrategyRollingUpdate `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#type Daemonset#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

