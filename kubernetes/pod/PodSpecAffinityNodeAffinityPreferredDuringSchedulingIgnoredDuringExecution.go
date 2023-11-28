// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// preference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/pod#preference Pod#preference}
	Preference *PodSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	// weight is in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/pod#weight Pod#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

