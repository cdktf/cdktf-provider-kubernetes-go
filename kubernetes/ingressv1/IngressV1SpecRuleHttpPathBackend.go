// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingressv1


type IngressV1SpecRuleHttpPathBackend struct {
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/ingress_v1#resource IngressV1#resource}
	Resource *IngressV1SpecRuleHttpPathBackendResource `field:"optional" json:"resource" yaml:"resource"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/ingress_v1#service IngressV1#service}
	Service *IngressV1SpecRuleHttpPathBackendService `field:"optional" json:"service" yaml:"service"`
}

