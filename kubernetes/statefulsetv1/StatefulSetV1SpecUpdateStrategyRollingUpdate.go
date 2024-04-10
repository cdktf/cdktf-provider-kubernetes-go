// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecUpdateStrategyRollingUpdate struct {
	// Indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/stateful_set_v1#partition StatefulSetV1#partition}
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
}

