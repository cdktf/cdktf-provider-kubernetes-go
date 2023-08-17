package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecMetricObject struct {
	// described_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/horizontal_pod_autoscaler_v2#described_object HorizontalPodAutoscalerV2#described_object}
	DescribedObject *HorizontalPodAutoscalerV2SpecMetricObjectDescribedObject `field:"required" json:"describedObject" yaml:"describedObject"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/horizontal_pod_autoscaler_v2#metric HorizontalPodAutoscalerV2#metric}
	Metric *HorizontalPodAutoscalerV2SpecMetricObjectMetric `field:"required" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/horizontal_pod_autoscaler_v2#target HorizontalPodAutoscalerV2#target}
	Target *HorizontalPodAutoscalerV2SpecMetricObjectTarget `field:"optional" json:"target" yaml:"target"`
}

