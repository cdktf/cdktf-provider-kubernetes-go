package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecScaleTargetRef struct {
	// Kind of the referent. e.g. `ReplicationController`. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/horizontal_pod_autoscaler#kind HorizontalPodAutoscaler#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/horizontal_pod_autoscaler#name HorizontalPodAutoscaler#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// API version of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/horizontal_pod_autoscaler#api_version HorizontalPodAutoscaler#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
}

