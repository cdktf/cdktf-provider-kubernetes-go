// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingress


type IngressSpec struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/ingress#backend Ingress#backend}
	Backend *IngressSpecBackend `field:"optional" json:"backend" yaml:"backend"`
	// ingressClassName is the name of the IngressClass cluster resource.
	//
	// The associated IngressClass defines which controller will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/ingress#ingress_class_name Ingress#ingress_class_name}
	IngressClassName *string `field:"optional" json:"ingressClassName" yaml:"ingressClassName"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/ingress#rule Ingress#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/ingress#tls Ingress#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

