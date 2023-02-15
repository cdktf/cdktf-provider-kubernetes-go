package deploymentv1


type DeploymentV1SpecStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#rolling_update DeploymentV1#rolling_update}
	RollingUpdate *DeploymentV1SpecStrategyRollingUpdate `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#type DeploymentV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

