// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manifest


type ManifestWaitCondition struct {
	// The condition status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/manifest#status Manifest#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The type of condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/manifest#type Manifest#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

