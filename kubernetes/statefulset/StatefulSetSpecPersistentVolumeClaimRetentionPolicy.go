// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecPersistentVolumeClaimRetentionPolicy struct {
	// This field controls what happens when a Statefulset is deleted. Default is Retain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#when_deleted StatefulSet#when_deleted}
	WhenDeleted *string `field:"optional" json:"whenDeleted" yaml:"whenDeleted"`
	// This field controls what happens when a Statefulset is scaled. Default is Retain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#when_scaled StatefulSet#when_scaled}
	WhenScaled *string `field:"optional" json:"whenScaled" yaml:"whenScaled"`
}

