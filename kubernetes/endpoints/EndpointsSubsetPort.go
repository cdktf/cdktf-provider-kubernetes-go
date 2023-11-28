// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpoints


type EndpointsSubsetPort struct {
	// The port that will be exposed by this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/endpoints#port Endpoints#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The name of this port within the endpoint.
	//
	// Must be a DNS_LABEL. Optional if only one Port is defined on this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/endpoints#name Endpoints#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IP protocol for this port. Supports `TCP` and `UDP`. Default is `TCP`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/endpoints#protocol Endpoints#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

