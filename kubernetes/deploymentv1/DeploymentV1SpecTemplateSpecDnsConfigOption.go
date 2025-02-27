// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1


type DeploymentV1SpecTemplateSpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/deployment_v1#name DeploymentV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/deployment_v1#value DeploymentV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

