// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterrolebindingv1


type ClusterRoleBindingV1Subject struct {
	// The kind of resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cluster_role_binding_v1#kind ClusterRoleBindingV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the resource to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cluster_role_binding_v1#name ClusterRoleBindingV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API group of the subject resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cluster_role_binding_v1#api_group ClusterRoleBindingV1#api_group}
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
	// The Namespace of the subject resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cluster_role_binding_v1#namespace ClusterRoleBindingV1#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

