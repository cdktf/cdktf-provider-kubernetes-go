// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscaler


type HorizontalPodAutoscalerSpec struct {
	// Upper limit for the number of pods that can be set by the autoscaler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler#max_replicas HorizontalPodAutoscaler#max_replicas}
	MaxReplicas *float64 `field:"required" json:"maxReplicas" yaml:"maxReplicas"`
	// scale_target_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler#scale_target_ref HorizontalPodAutoscaler#scale_target_ref}
	ScaleTargetRef *HorizontalPodAutoscalerSpecScaleTargetRef `field:"required" json:"scaleTargetRef" yaml:"scaleTargetRef"`
	// behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler#behavior HorizontalPodAutoscaler#behavior}
	Behavior *HorizontalPodAutoscalerSpecBehavior `field:"optional" json:"behavior" yaml:"behavior"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler#metric HorizontalPodAutoscaler#metric}
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// Lower limit for the number of pods that can be set by the autoscaler, defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler#min_replicas HorizontalPodAutoscaler#min_replicas}
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// Target average CPU utilization (represented as a percentage of requested CPU) over all the pods.
	//
	// If not specified the default autoscaling policy will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/horizontal_pod_autoscaler#target_cpu_utilization_percentage HorizontalPodAutoscaler#target_cpu_utilization_percentage}
	TargetCpuUtilizationPercentage *float64 `field:"optional" json:"targetCpuUtilizationPercentage" yaml:"targetCpuUtilizationPercentage"`
}

