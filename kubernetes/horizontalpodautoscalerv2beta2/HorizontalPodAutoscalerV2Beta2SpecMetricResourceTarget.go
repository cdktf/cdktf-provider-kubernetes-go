// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2beta2


type HorizontalPodAutoscalerV2Beta2SpecMetricResourceTarget struct {
	// type represents whether the metric type is Utilization, Value, or AverageValue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler_v2beta2#type HorizontalPodAutoscalerV2Beta2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
	//
	// Currently only valid for Resource metric source type
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler_v2beta2#average_utilization HorizontalPodAutoscalerV2Beta2#average_utilization}
	AverageUtilization *float64 `field:"optional" json:"averageUtilization" yaml:"averageUtilization"`
	// averageValue is the target value of the average of the metric across all relevant pods (as a quantity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler_v2beta2#average_value HorizontalPodAutoscalerV2Beta2#average_value}
	AverageValue *string `field:"optional" json:"averageValue" yaml:"averageValue"`
	// value is the target value of the metric (as a quantity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler_v2beta2#value HorizontalPodAutoscalerV2Beta2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

