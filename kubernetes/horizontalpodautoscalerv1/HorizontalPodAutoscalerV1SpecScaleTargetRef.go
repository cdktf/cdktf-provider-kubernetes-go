// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package horizontalpodautoscalerv1


type HorizontalPodAutoscalerV1SpecScaleTargetRef struct {
	// Kind of the referent. e.g. `ReplicationController`. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/horizontal_pod_autoscaler_v1#kind HorizontalPodAutoscalerV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/horizontal_pod_autoscaler_v1#name HorizontalPodAutoscalerV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// API version of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/horizontal_pod_autoscaler_v1#api_version HorizontalPodAutoscalerV1#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

