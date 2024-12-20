// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingressv1


type IngressV1SpecDefaultBackend struct {
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/ingress_v1#resource IngressV1#resource}
	Resource *IngressV1SpecDefaultBackendResource `field:"optional" json:"resource" yaml:"resource"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/ingress_v1#service IngressV1#service}
	Service *IngressV1SpecDefaultBackendService `field:"optional" json:"service" yaml:"service"`
}

