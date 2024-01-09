// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesendpointsv1


type DataKubernetesEndpointsV1SubsetPort struct {
	// The port that will be exposed by this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/endpoints_v1#port DataKubernetesEndpointsV1#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The name of this port within the endpoint.
	//
	// Must be a DNS_LABEL. Optional if only one Port is defined on this endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/endpoints_v1#name DataKubernetesEndpointsV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IP protocol for this port. Supports `TCP` and `UDP`. Default is `TCP`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/endpoints_v1#protocol DataKubernetesEndpointsV1#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

