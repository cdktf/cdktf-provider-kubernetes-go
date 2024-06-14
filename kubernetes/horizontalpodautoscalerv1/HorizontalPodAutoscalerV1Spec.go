// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv1


type HorizontalPodAutoscalerV1Spec struct {
	// Upper limit for the number of pods that can be set by the autoscaler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/horizontal_pod_autoscaler_v1#max_replicas HorizontalPodAutoscalerV1#max_replicas}
	MaxReplicas *float64 `field:"required" json:"maxReplicas" yaml:"maxReplicas"`
	// scale_target_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/horizontal_pod_autoscaler_v1#scale_target_ref HorizontalPodAutoscalerV1#scale_target_ref}
	ScaleTargetRef *HorizontalPodAutoscalerV1SpecScaleTargetRef `field:"required" json:"scaleTargetRef" yaml:"scaleTargetRef"`
	// Lower limit for the number of pods that can be set by the autoscaler, defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/horizontal_pod_autoscaler_v1#min_replicas HorizontalPodAutoscalerV1#min_replicas}
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// Target average CPU utilization (represented as a percentage of requested CPU) over all the pods.
	//
	// If not specified the default autoscaling policy will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/horizontal_pod_autoscaler_v1#target_cpu_utilization_percentage HorizontalPodAutoscalerV1#target_cpu_utilization_percentage}
	TargetCpuUtilizationPercentage *float64 `field:"optional" json:"targetCpuUtilizationPercentage" yaml:"targetCpuUtilizationPercentage"`
}

