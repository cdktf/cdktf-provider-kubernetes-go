// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingress


type IngressSpecRuleHttpPathBackend struct {
	// Specifies the name of the referenced service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/ingress#service_name Ingress#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Specifies the port of the referenced service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/ingress#service_port Ingress#service_port}
	ServicePort *string `field:"optional" json:"servicePort" yaml:"servicePort"`
}

