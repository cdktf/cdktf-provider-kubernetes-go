// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package annotations


type AnnotationsMetadata struct {
	// The name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/annotations#name Annotations#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/annotations#namespace Annotations#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

