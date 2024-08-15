// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	// empty topology key is interpreted by the scheduler as 'all topologies'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemonset#topology_key Daemonset#topology_key}
	TopologyKey *string `field:"required" json:"topologyKey" yaml:"topologyKey"`
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemonset#label_selector Daemonset#label_selector}
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemonset#namespaces Daemonset#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

