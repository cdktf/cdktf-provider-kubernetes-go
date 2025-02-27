// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolume


type PersistentVolumeSpecNodeAffinityRequired struct {
	// node_selector_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/persistent_volume#node_selector_term PersistentVolume#node_selector_term}
	NodeSelectorTerm interface{} `field:"required" json:"nodeSelectorTerm" yaml:"nodeSelectorTerm"`
}

