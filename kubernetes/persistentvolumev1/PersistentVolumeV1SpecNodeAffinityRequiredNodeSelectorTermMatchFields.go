// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumev1


type PersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTermMatchFields struct {
	// The label key that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/persistent_volume_v1#key PersistentVolumeV1#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists`, `DoesNotExist`, `Gt`, and `Lt`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/persistent_volume_v1#operator PersistentVolumeV1#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// An array of string values.
	//
	// If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/persistent_volume_v1#values PersistentVolumeV1#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

