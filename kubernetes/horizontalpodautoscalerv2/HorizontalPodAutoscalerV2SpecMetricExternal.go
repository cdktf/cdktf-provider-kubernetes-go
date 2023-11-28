// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecMetricExternal struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler_v2#metric HorizontalPodAutoscalerV2#metric}
	Metric *HorizontalPodAutoscalerV2SpecMetricExternalMetric `field:"required" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler_v2#target HorizontalPodAutoscalerV2#target}
	Target *HorizontalPodAutoscalerV2SpecMetricExternalTarget `field:"optional" json:"target" yaml:"target"`
}

