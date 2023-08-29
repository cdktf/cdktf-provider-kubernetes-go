// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterrolev1


type ClusterRoleV1Rule struct {
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.
	//
	// VerbAll represents all kinds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cluster_role_v1#verbs ClusterRoleV1#verbs}
	Verbs *[]*string `field:"required" json:"verbs" yaml:"verbs"`
	// APIGroups is the name of the APIGroup that contains the resources.
	//
	// If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cluster_role_v1#api_groups ClusterRoleV1#api_groups}
	ApiGroups *[]*string `field:"optional" json:"apiGroups" yaml:"apiGroups"`
	// NonResourceURLs is a set of partial urls that a user should have access to.
	//
	// *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"), but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cluster_role_v1#non_resource_urls ClusterRoleV1#non_resource_urls}
	NonResourceUrls *[]*string `field:"optional" json:"nonResourceUrls" yaml:"nonResourceUrls"`
	// ResourceNames is an optional white list of names that the rule applies to.
	//
	// An empty set means that everything is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cluster_role_v1#resource_names ClusterRoleV1#resource_names}
	ResourceNames *[]*string `field:"optional" json:"resourceNames" yaml:"resourceNames"`
	// Resources is a list of resources this rule applies to. ResourceAll represents all resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cluster_role_v1#resources ClusterRoleV1#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

