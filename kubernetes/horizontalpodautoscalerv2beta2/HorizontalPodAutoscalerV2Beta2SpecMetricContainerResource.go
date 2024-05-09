// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2beta2


type HorizontalPodAutoscalerV2Beta2SpecMetricContainerResource struct {
	// name of the container in the pods of the scaling target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/horizontal_pod_autoscaler_v2beta2#container HorizontalPodAutoscalerV2Beta2#container}
	Container *string `field:"required" json:"container" yaml:"container"`
	// name of the resource in question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/horizontal_pod_autoscaler_v2beta2#name HorizontalPodAutoscalerV2Beta2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/horizontal_pod_autoscaler_v2beta2#target HorizontalPodAutoscalerV2Beta2#target}
	Target *HorizontalPodAutoscalerV2Beta2SpecMetricContainerResourceTarget `field:"optional" json:"target" yaml:"target"`
}

