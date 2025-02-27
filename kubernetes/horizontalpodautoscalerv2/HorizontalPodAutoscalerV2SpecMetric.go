// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecMetric struct {
	// type is the type of metric source.
	//
	// It should be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each mapping to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate HPAContainerMetrics is enabled
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/horizontal_pod_autoscaler_v2#type HorizontalPodAutoscalerV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// container_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/horizontal_pod_autoscaler_v2#container_resource HorizontalPodAutoscalerV2#container_resource}
	ContainerResource *HorizontalPodAutoscalerV2SpecMetricContainerResource `field:"optional" json:"containerResource" yaml:"containerResource"`
	// external block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/horizontal_pod_autoscaler_v2#external HorizontalPodAutoscalerV2#external}
	External *HorizontalPodAutoscalerV2SpecMetricExternal `field:"optional" json:"external" yaml:"external"`
	// object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/horizontal_pod_autoscaler_v2#object HorizontalPodAutoscalerV2#object}
	Object *HorizontalPodAutoscalerV2SpecMetricObject `field:"optional" json:"object" yaml:"object"`
	// pods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/horizontal_pod_autoscaler_v2#pods HorizontalPodAutoscalerV2#pods}
	Pods *HorizontalPodAutoscalerV2SpecMetricPods `field:"optional" json:"pods" yaml:"pods"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/horizontal_pod_autoscaler_v2#resource HorizontalPodAutoscalerV2#resource}
	Resource *HorizontalPodAutoscalerV2SpecMetricResource `field:"optional" json:"resource" yaml:"resource"`
}

