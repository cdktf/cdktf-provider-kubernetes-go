// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	// empty topology key is interpreted by the scheduler as 'all topologies'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#topology_key Job#topology_key}
	TopologyKey *string `field:"required" json:"topologyKey" yaml:"topologyKey"`
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#label_selector Job#label_selector}
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#namespaces Job#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// namespace_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#namespace_selector Job#namespace_selector}
	NamespaceSelector interface{} `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
}

