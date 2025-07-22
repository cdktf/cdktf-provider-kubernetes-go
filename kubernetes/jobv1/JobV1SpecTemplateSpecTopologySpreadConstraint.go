// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecTemplateSpecTopologySpreadConstraint struct {
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#label_selector JobV1#label_selector}
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// is a set of pod label keys to select the pods over which spreading will be calculated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#match_label_keys JobV1#match_label_keys}
	MatchLabelKeys *[]*string `field:"optional" json:"matchLabelKeys" yaml:"matchLabelKeys"`
	// describes the degree to which pods may be unevenly distributed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#max_skew JobV1#max_skew}
	MaxSkew *float64 `field:"optional" json:"maxSkew" yaml:"maxSkew"`
	// indicates a minimum number of eligible domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#min_domains JobV1#min_domains}
	MinDomains *float64 `field:"optional" json:"minDomains" yaml:"minDomains"`
	// indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#node_affinity_policy JobV1#node_affinity_policy}
	NodeAffinityPolicy *string `field:"optional" json:"nodeAffinityPolicy" yaml:"nodeAffinityPolicy"`
	// indicates how we will treat node taints when calculating pod topology spread skew.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#node_taints_policy JobV1#node_taints_policy}
	NodeTaintsPolicy *string `field:"optional" json:"nodeTaintsPolicy" yaml:"nodeTaintsPolicy"`
	// the key of node labels.
	//
	// Nodes that have a label with this key and identical values are considered to be in the same topology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#topology_key JobV1#topology_key}
	TopologyKey *string `field:"optional" json:"topologyKey" yaml:"topologyKey"`
	// indicates how to deal with a pod if it doesn't satisfy the spread constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/job_v1#when_unsatisfiable JobV1#when_unsatisfiable}
	WhenUnsatisfiable *string `field:"optional" json:"whenUnsatisfiable" yaml:"whenUnsatisfiable"`
}

