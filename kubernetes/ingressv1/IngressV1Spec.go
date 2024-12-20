// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingressv1


type IngressV1Spec struct {
	// default_backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/ingress_v1#default_backend IngressV1#default_backend}
	DefaultBackend *IngressV1SpecDefaultBackend `field:"optional" json:"defaultBackend" yaml:"defaultBackend"`
	// ingressClassName is the name of an IngressClass cluster resource.
	//
	// Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the `kubernetes.io/ingress.class` annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/ingress_v1#ingress_class_name IngressV1#ingress_class_name}
	IngressClassName *string `field:"optional" json:"ingressClassName" yaml:"ingressClassName"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/ingress_v1#rule IngressV1#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/ingress_v1#tls IngressV1#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

