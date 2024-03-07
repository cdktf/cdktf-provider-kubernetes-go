// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manifest


type ManifestTimeouts struct {
	// Timeout for the create operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/manifest#create Manifest#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Timeout for the delete operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/manifest#delete Manifest#delete}
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Timeout for the update operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/manifest#update Manifest#update}
	Update *string `field:"optional" json:"update" yaml:"update"`
}

