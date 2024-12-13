// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2beta2


type HorizontalPodAutoscalerV2Beta2SpecMetricPods struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/horizontal_pod_autoscaler_v2beta2#metric HorizontalPodAutoscalerV2Beta2#metric}
	Metric *HorizontalPodAutoscalerV2Beta2SpecMetricPodsMetric `field:"required" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/horizontal_pod_autoscaler_v2beta2#target HorizontalPodAutoscalerV2Beta2#target}
	Target *HorizontalPodAutoscalerV2Beta2SpecMetricPodsTarget `field:"optional" json:"target" yaml:"target"`
}

