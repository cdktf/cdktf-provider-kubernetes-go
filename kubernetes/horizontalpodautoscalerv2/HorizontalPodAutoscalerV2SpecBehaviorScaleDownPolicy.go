// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecBehaviorScaleDownPolicy struct {
	// Period specifies the window of time for which the policy should hold true.
	//
	// PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/horizontal_pod_autoscaler_v2#period_seconds HorizontalPodAutoscalerV2#period_seconds}
	PeriodSeconds *float64 `field:"required" json:"periodSeconds" yaml:"periodSeconds"`
	// Type is used to specify the scaling policy: Percent or Pods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/horizontal_pod_autoscaler_v2#type HorizontalPodAutoscalerV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value contains the amount of change which is permitted by the policy. It must be greater than zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/horizontal_pod_autoscaler_v2#value HorizontalPodAutoscalerV2#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

