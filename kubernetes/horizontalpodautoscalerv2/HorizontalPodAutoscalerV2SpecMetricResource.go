// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecMetricResource struct {
	// name is the name of the resource in question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/horizontal_pod_autoscaler_v2#name HorizontalPodAutoscalerV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/horizontal_pod_autoscaler_v2#target HorizontalPodAutoscalerV2#target}
	Target *HorizontalPodAutoscalerV2SpecMetricResourceTarget `field:"optional" json:"target" yaml:"target"`
}

