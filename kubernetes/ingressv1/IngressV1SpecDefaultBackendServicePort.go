// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingressv1


type IngressV1SpecDefaultBackendServicePort struct {
	// Specifies the name of the port of the referenced service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_v1#name IngressV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the numerical port of the referenced service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_v1#number IngressV1#number}
	Number *float64 `field:"optional" json:"number" yaml:"number"`
}

