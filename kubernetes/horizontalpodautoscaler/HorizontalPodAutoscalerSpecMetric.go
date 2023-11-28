// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecMetric struct {
	// type is the type of metric source.
	//
	// It should be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each mapping to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate HPAContainerMetrics is enabled
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#type HorizontalPodAutoscaler#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// container_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#container_resource HorizontalPodAutoscaler#container_resource}
	ContainerResource *HorizontalPodAutoscalerSpecMetricContainerResource `field:"optional" json:"containerResource" yaml:"containerResource"`
	// external block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#external HorizontalPodAutoscaler#external}
	External *HorizontalPodAutoscalerSpecMetricExternal `field:"optional" json:"external" yaml:"external"`
	// object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#object HorizontalPodAutoscaler#object}
	Object *HorizontalPodAutoscalerSpecMetricObject `field:"optional" json:"object" yaml:"object"`
	// pods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#pods HorizontalPodAutoscaler#pods}
	Pods *HorizontalPodAutoscalerSpecMetricPods `field:"optional" json:"pods" yaml:"pods"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#resource HorizontalPodAutoscaler#resource}
	Resource *HorizontalPodAutoscalerSpecMetricResource `field:"optional" json:"resource" yaml:"resource"`
}

