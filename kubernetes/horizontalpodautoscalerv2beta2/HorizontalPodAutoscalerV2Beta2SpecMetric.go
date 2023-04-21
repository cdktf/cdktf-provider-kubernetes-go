package horizontalpodautoscalerv2beta2


type HorizontalPodAutoscalerV2Beta2SpecMetric struct {
	// type is the type of metric source.
	//
	// It should be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each mapping to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate HPAContainerMetrics is enabled
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler_v2beta2#type HorizontalPodAutoscalerV2Beta2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// container_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler_v2beta2#container_resource HorizontalPodAutoscalerV2Beta2#container_resource}
	ContainerResource *HorizontalPodAutoscalerV2Beta2SpecMetricContainerResource `field:"optional" json:"containerResource" yaml:"containerResource"`
	// external block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler_v2beta2#external HorizontalPodAutoscalerV2Beta2#external}
	External *HorizontalPodAutoscalerV2Beta2SpecMetricExternal `field:"optional" json:"external" yaml:"external"`
	// object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler_v2beta2#object HorizontalPodAutoscalerV2Beta2#object}
	Object *HorizontalPodAutoscalerV2Beta2SpecMetricObject `field:"optional" json:"object" yaml:"object"`
	// pods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler_v2beta2#pods HorizontalPodAutoscalerV2Beta2#pods}
	Pods *HorizontalPodAutoscalerV2Beta2SpecMetricPods `field:"optional" json:"pods" yaml:"pods"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler_v2beta2#resource HorizontalPodAutoscalerV2Beta2#resource}
	Resource *HorizontalPodAutoscalerV2Beta2SpecMetricResource `field:"optional" json:"resource" yaml:"resource"`
}

