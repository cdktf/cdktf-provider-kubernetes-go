// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecMetricResourceTarget struct {
	// type represents whether the metric type is Utilization, Value, or AverageValue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/horizontal_pod_autoscaler_v2#type HorizontalPodAutoscalerV2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
	//
	// Currently only valid for Resource metric source type
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/horizontal_pod_autoscaler_v2#average_utilization HorizontalPodAutoscalerV2#average_utilization}
	AverageUtilization *float64 `field:"optional" json:"averageUtilization" yaml:"averageUtilization"`
	// averageValue is the target value of the average of the metric across all relevant pods (as a quantity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/horizontal_pod_autoscaler_v2#average_value HorizontalPodAutoscalerV2#average_value}
	AverageValue *string `field:"optional" json:"averageValue" yaml:"averageValue"`
	// value is the target value of the metric (as a quantity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/horizontal_pod_autoscaler_v2#value HorizontalPodAutoscalerV2#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

