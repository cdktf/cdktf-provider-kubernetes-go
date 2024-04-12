// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterrolebinding


type ClusterRoleBindingSubject struct {
	// The kind of resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/cluster_role_binding#kind ClusterRoleBinding#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The name of the resource to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/cluster_role_binding#name ClusterRoleBinding#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API group of the subject resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/cluster_role_binding#api_group ClusterRoleBinding#api_group}
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
	// The Namespace of the subject resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/cluster_role_binding#namespace ClusterRoleBinding#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

