// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nodetaint


type NodeTaintMetadata struct {
	// The name of the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/node_taint#name NodeTaint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

