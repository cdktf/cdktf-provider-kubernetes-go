// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontrollerv1


type ReplicationControllerV1Spec struct {
	// A label query over pods that should match the Replicas count.
	//
	// If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#selector ReplicationControllerV1#selector}
	Selector *map[string]*string `field:"required" json:"selector" yaml:"selector"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#template ReplicationControllerV1#template}
	Template *ReplicationControllerV1SpecTemplate `field:"required" json:"template" yaml:"template"`
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	//
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#min_ready_seconds ReplicationControllerV1#min_ready_seconds}
	MinReadySeconds *float64 `field:"optional" json:"minReadySeconds" yaml:"minReadySeconds"`
	// The number of desired replicas. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#replicas ReplicationControllerV1#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

