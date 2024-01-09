// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecTopologySpreadConstraintLabelSelectorMatchExpressions struct {
	// The label key that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#key PodV1#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#operator PodV1#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// An array of string values.
	//
	// If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_v1#values PodV1#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

