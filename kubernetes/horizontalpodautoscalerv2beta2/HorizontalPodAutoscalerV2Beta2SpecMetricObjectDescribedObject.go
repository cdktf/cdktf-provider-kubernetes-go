// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv2beta2


type HorizontalPodAutoscalerV2Beta2SpecMetricObjectDescribedObject struct {
	// API version of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/horizontal_pod_autoscaler_v2beta2#api_version HorizontalPodAutoscalerV2Beta2#api_version}
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/horizontal_pod_autoscaler_v2beta2#kind HorizontalPodAutoscalerV2Beta2#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/horizontal_pod_autoscaler_v2beta2#name HorizontalPodAutoscalerV2Beta2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

