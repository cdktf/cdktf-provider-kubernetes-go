// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecPersistentVolumeClaimRetentionPolicy struct {
	// This field controls what happens when a Statefulset is deleted. Default is Retain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set_v1#when_deleted StatefulSetV1#when_deleted}
	WhenDeleted *string `field:"optional" json:"whenDeleted" yaml:"whenDeleted"`
	// This field controls what happens when a Statefulset is scaled. Default is Retain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set_v1#when_scaled StatefulSetV1#when_scaled}
	WhenScaled *string `field:"optional" json:"whenScaled" yaml:"whenScaled"`
}

