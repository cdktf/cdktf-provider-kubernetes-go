// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate struct {
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/deployment#spec Deployment#spec}
	Spec *DeploymentSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/deployment#metadata Deployment#metadata}
	Metadata *DeploymentSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

