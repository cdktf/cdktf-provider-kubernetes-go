// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecTopologySpreadConstraintLabelSelector struct {
	// match_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#match_expressions PodV1#match_expressions}
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// A map of {key,value} pairs.
	//
	// A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#match_labels PodV1#match_labels}
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

