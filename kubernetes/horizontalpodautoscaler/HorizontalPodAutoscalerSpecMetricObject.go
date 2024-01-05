// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecMetricObject struct {
	// described_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/horizontal_pod_autoscaler#described_object HorizontalPodAutoscaler#described_object}
	DescribedObject *HorizontalPodAutoscalerSpecMetricObjectDescribedObject `field:"required" json:"describedObject" yaml:"describedObject"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/horizontal_pod_autoscaler#metric HorizontalPodAutoscaler#metric}
	Metric *HorizontalPodAutoscalerSpecMetricObjectMetric `field:"required" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/horizontal_pod_autoscaler#target HorizontalPodAutoscaler#target}
	Target *HorizontalPodAutoscalerSpecMetricObjectTarget `field:"optional" json:"target" yaml:"target"`
}

