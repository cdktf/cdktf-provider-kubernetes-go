// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2beta2


type HorizontalPodAutoscalerV2Beta2SpecMetricPodsMetric struct {
	// name is the name of the given metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/horizontal_pod_autoscaler_v2beta2#name HorizontalPodAutoscalerV2Beta2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/horizontal_pod_autoscaler_v2beta2#selector HorizontalPodAutoscalerV2Beta2#selector}
	Selector interface{} `field:"optional" json:"selector" yaml:"selector"`
}

