package deployment


type DeploymentSpecStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#rolling_update Deployment#rolling_update}
	RollingUpdate *DeploymentSpecStrategyRollingUpdate `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#type Deployment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

