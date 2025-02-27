// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecBehavior struct {
	// scale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/horizontal_pod_autoscaler_v2#scale_down HorizontalPodAutoscalerV2#scale_down}
	ScaleDown interface{} `field:"optional" json:"scaleDown" yaml:"scaleDown"`
	// scale_up block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/horizontal_pod_autoscaler_v2#scale_up HorizontalPodAutoscalerV2#scale_up}
	ScaleUp interface{} `field:"optional" json:"scaleUp" yaml:"scaleUp"`
}

