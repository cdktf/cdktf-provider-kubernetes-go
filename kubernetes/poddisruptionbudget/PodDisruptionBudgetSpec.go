// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package poddisruptionbudget


type PodDisruptionBudgetSpec struct {
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/pod_disruption_budget#selector PodDisruptionBudget#selector}
	Selector *PodDisruptionBudgetSpecSelector `field:"required" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/pod_disruption_budget#max_unavailable PodDisruptionBudget#max_unavailable}.
	MaxUnavailable *string `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/pod_disruption_budget#min_available PodDisruptionBudget#min_available}.
	MinAvailable *string `field:"optional" json:"minAvailable" yaml:"minAvailable"`
}

