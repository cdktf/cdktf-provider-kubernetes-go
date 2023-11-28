// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1


type DeploymentV1SpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment_v1#metadata DeploymentV1#metadata}
	Metadata *DeploymentV1SpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment_v1#spec DeploymentV1#spec}
	Spec *DeploymentV1SpecTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

