// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolebindingv1


type RoleBindingV1Subject struct {
	// The kind of resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/role_binding_v1#kind RoleBindingV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the resource to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/role_binding_v1#name RoleBindingV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API group of the subject resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/role_binding_v1#api_group RoleBindingV1#api_group}
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
	// The Namespace of the subject resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/role_binding_v1#namespace RoleBindingV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

