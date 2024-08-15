// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	// node_selector_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/pod#node_selector_term Pod#node_selector_term}
	NodeSelectorTerm interface{} `field:"optional" json:"nodeSelectorTerm" yaml:"nodeSelectorTerm"`
}

