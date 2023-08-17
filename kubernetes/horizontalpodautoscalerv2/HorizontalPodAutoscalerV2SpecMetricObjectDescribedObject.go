package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecMetricObjectDescribedObject struct {
	// API version of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/horizontal_pod_autoscaler_v2#api_version HorizontalPodAutoscalerV2#api_version}
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/horizontal_pod_autoscaler_v2#kind HorizontalPodAutoscalerV2#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/horizontal_pod_autoscaler_v2#name HorizontalPodAutoscalerV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

