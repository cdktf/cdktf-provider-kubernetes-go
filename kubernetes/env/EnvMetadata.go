// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package env


type EnvMetadata struct {
	// The name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/env#name Env#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/env#namespace Env#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

