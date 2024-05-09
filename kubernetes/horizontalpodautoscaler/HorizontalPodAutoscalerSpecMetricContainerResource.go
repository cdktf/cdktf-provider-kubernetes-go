// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecMetricContainerResource struct {
	// name of the container in the pods of the scaling target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/horizontal_pod_autoscaler#container HorizontalPodAutoscaler#container}
	Container *string `field:"required" json:"container" yaml:"container"`
	// name of the resource in question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/horizontal_pod_autoscaler#name HorizontalPodAutoscaler#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/horizontal_pod_autoscaler#target HorizontalPodAutoscaler#target}
	Target *HorizontalPodAutoscalerSpecMetricContainerResourceTarget `field:"optional" json:"target" yaml:"target"`
}

