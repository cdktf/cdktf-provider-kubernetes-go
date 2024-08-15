// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterrolev1


type ClusterRoleV1AggregationRule struct {
	// cluster_role_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cluster_role_v1#cluster_role_selectors ClusterRoleV1#cluster_role_selectors}
	ClusterRoleSelectors interface{} `field:"optional" json:"clusterRoleSelectors" yaml:"clusterRoleSelectors"`
}

