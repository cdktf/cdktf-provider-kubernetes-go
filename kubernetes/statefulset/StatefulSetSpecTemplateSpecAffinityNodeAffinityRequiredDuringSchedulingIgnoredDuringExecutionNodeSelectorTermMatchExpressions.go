// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermMatchExpressions struct {
	// The label key that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/stateful_set#key StatefulSet#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Operator represents a key's relationship to a set of values.
	//
	// Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/stateful_set#operator StatefulSet#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Values is an array of string values.
	//
	// If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/stateful_set#values StatefulSet#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

