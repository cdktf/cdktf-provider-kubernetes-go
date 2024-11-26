// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labels


type LabelsMetadata struct {
	// The name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/labels#name Labels#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/labels#namespace Labels#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

