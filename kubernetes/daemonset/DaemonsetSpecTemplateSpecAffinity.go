// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecAffinity struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemonset#node_affinity Daemonset#node_affinity}
	NodeAffinity *DaemonsetSpecTemplateSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// pod_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemonset#pod_affinity Daemonset#pod_affinity}
	PodAffinity *DaemonsetSpecTemplateSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_anti_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemonset#pod_anti_affinity Daemonset#pod_anti_affinity}
	PodAntiAffinity *DaemonsetSpecTemplateSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

