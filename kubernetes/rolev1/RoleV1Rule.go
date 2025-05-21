// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolev1


type RoleV1Rule struct {
	// Name of the APIGroup that contains the resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/role_v1#api_groups RoleV1#api_groups}
	ApiGroups *[]*string `field:"required" json:"apiGroups" yaml:"apiGroups"`
	// List of resources that the rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/role_v1#resources RoleV1#resources}
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
	// List of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/role_v1#verbs RoleV1#verbs}
	Verbs *[]*string `field:"required" json:"verbs" yaml:"verbs"`
	// White list of names that the rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/role_v1#resource_names RoleV1#resource_names}
	ResourceNames *[]*string `field:"optional" json:"resourceNames" yaml:"resourceNames"`
}

