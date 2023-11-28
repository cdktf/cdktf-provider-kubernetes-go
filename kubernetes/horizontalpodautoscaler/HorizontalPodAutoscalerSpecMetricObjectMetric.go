// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecMetricObjectMetric struct {
	// name is the name of the given metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#name HorizontalPodAutoscaler#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#selector HorizontalPodAutoscaler#selector}
	Selector interface{} `field:"optional" json:"selector" yaml:"selector"`
}

