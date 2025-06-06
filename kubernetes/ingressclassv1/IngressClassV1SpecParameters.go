// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingressclassv1


type IngressClassV1SpecParameters struct {
	// Kind is the type of resource being referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_class_v1#kind IngressClassV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name is the name of resource being referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_class_v1#name IngressClassV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// APIGroup is the group for the resource being referenced.
	//
	// If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_class_v1#api_group IngressClassV1#api_group}
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_class_v1#namespace IngressClassV1#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_class_v1#scope IngressClassV1#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

