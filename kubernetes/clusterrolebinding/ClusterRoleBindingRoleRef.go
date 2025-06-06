// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterrolebinding


type ClusterRoleBindingRoleRef struct {
	// The API group of the user. The only value possible at the moment is `rbac.authorization.k8s.io`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cluster_role_binding#api_group ClusterRoleBinding#api_group}
	ApiGroup *string `field:"required" json:"apiGroup" yaml:"apiGroup"`
	// The kind of resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cluster_role_binding#kind ClusterRoleBinding#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the User to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cluster_role_binding#name ClusterRoleBinding#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

