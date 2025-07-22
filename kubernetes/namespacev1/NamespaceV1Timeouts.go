// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namespacev1


type NamespaceV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/namespace_v1#delete NamespaceV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

