// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecSecurityContextSeLinuxOptions struct {
	// Level is SELinux level label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/deployment#level Deployment#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Role is a SELinux role label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/deployment#role Deployment#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Type is a SELinux type label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/deployment#type Deployment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User is a SELinux user label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/deployment#user Deployment#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

