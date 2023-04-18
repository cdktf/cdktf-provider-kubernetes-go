package horizontalpodautoscalerv1


type HorizontalPodAutoscalerV1SpecScaleTargetRef struct {
	// Kind of the referent. e.g. `ReplicationController`. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/horizontal_pod_autoscaler_v1#kind HorizontalPodAutoscalerV1#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/horizontal_pod_autoscaler_v1#name HorizontalPodAutoscalerV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// API version of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/horizontal_pod_autoscaler_v1#api_version HorizontalPodAutoscalerV1#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

