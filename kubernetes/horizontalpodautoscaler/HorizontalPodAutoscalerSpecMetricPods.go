// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecMetricPods struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/horizontal_pod_autoscaler#metric HorizontalPodAutoscaler#metric}
	Metric *HorizontalPodAutoscalerSpecMetricPodsMetric `field:"required" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/horizontal_pod_autoscaler#target HorizontalPodAutoscaler#target}
	Target *HorizontalPodAutoscalerSpecMetricPodsTarget `field:"optional" json:"target" yaml:"target"`
}

