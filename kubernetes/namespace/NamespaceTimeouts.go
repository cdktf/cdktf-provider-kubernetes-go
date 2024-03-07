// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namespace


type NamespaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/namespace#delete Namespace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

