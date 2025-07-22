// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpec struct {
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#template Daemonset#template}
	Template *DaemonsetSpecTemplate `field:"required" json:"template" yaml:"template"`
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	//
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#min_ready_seconds Daemonset#min_ready_seconds}
	MinReadySeconds *float64 `field:"optional" json:"minReadySeconds" yaml:"minReadySeconds"`
	// The number of old ReplicaSets to retain to allow rollback.
	//
	// This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#revision_history_limit Daemonset#revision_history_limit}
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#selector Daemonset#selector}
	Selector *DaemonsetSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#strategy Daemonset#strategy}
	Strategy *DaemonsetSpecStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

