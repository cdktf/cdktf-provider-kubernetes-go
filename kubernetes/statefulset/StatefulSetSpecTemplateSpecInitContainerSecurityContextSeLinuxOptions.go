// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecInitContainerSecurityContextSeLinuxOptions struct {
	// Level is SELinux level label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set#level StatefulSet#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Role is a SELinux role label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set#role StatefulSet#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Type is a SELinux type label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set#type StatefulSet#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User is a SELinux user label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set#user StatefulSet#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

