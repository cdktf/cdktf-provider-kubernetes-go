// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecContainerEnvFromSecretRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/deployment#name Deployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the Secret must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/deployment#optional Deployment#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

