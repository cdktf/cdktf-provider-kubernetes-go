// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecTopologySpreadConstraint struct {
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/pod#label_selector Pod#label_selector}
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// is a set of pod label keys to select the pods over which spreading will be calculated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/pod#match_label_keys Pod#match_label_keys}
	MatchLabelKeys *[]*string `field:"optional" json:"matchLabelKeys" yaml:"matchLabelKeys"`
	// describes the degree to which pods may be unevenly distributed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/pod#max_skew Pod#max_skew}
	MaxSkew *float64 `field:"optional" json:"maxSkew" yaml:"maxSkew"`
	// indicates a minimum number of eligible domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/pod#min_domains Pod#min_domains}
	MinDomains *float64 `field:"optional" json:"minDomains" yaml:"minDomains"`
	// indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/pod#node_affinity_policy Pod#node_affinity_policy}
	NodeAffinityPolicy *string `field:"optional" json:"nodeAffinityPolicy" yaml:"nodeAffinityPolicy"`
	// indicates how we will treat node taints when calculating pod topology spread skew.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/pod#node_taints_policy Pod#node_taints_policy}
	NodeTaintsPolicy *string `field:"optional" json:"nodeTaintsPolicy" yaml:"nodeTaintsPolicy"`
	// the key of node labels.
	//
	// Nodes that have a label with this key and identical values are considered to be in the same topology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/pod#topology_key Pod#topology_key}
	TopologyKey *string `field:"optional" json:"topologyKey" yaml:"topologyKey"`
	// indicates how to deal with a pod if it doesn't satisfy the spread constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/pod#when_unsatisfiable Pod#when_unsatisfiable}
	WhenUnsatisfiable *string `field:"optional" json:"whenUnsatisfiable" yaml:"whenUnsatisfiable"`
}

