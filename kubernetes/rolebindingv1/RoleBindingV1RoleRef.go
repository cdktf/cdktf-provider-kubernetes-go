// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolebindingv1


type RoleBindingV1RoleRef struct {
	// The API group of the user. The only value possible at the moment is `rbac.authorization.k8s.io`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/role_binding_v1#api_group RoleBindingV1#api_group}
	ApiGroup *string `field:"required" json:"apiGroup" yaml:"apiGroup"`
	// The kind of resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/role_binding_v1#kind RoleBindingV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the User to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/role_binding_v1#name RoleBindingV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

