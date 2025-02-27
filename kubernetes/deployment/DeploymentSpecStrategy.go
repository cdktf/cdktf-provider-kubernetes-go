// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/deployment#rolling_update Deployment#rolling_update}
	RollingUpdate *DeploymentSpecStrategyRollingUpdate `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/deployment#type Deployment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

