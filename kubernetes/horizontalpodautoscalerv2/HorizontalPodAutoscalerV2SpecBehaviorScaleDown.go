// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecBehaviorScaleDown struct {
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/horizontal_pod_autoscaler_v2#policy HorizontalPodAutoscalerV2#policy}
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// Used to specify which policy should be used. If not set, the default value Max is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/horizontal_pod_autoscaler_v2#select_policy HorizontalPodAutoscalerV2#select_policy}
	SelectPolicy *string `field:"optional" json:"selectPolicy" yaml:"selectPolicy"`
	// Number of seconds for which past recommendations should be considered while scaling up or scaling down.
	//
	// This value must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/horizontal_pod_autoscaler_v2#stabilization_window_seconds HorizontalPodAutoscalerV2#stabilization_window_seconds}
	StabilizationWindowSeconds *float64 `field:"optional" json:"stabilizationWindowSeconds" yaml:"stabilizationWindowSeconds"`
}

