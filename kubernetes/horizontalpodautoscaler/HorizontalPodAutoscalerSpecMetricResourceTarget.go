// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecMetricResourceTarget struct {
	// type represents whether the metric type is Utilization, Value, or AverageValue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#type HorizontalPodAutoscaler#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
	//
	// Currently only valid for Resource metric source type
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#average_utilization HorizontalPodAutoscaler#average_utilization}
	AverageUtilization *float64 `field:"optional" json:"averageUtilization" yaml:"averageUtilization"`
	// averageValue is the target value of the average of the metric across all relevant pods (as a quantity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#average_value HorizontalPodAutoscaler#average_value}
	AverageValue *string `field:"optional" json:"averageValue" yaml:"averageValue"`
	// value is the target value of the metric (as a quantity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/horizontal_pod_autoscaler#value HorizontalPodAutoscaler#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

