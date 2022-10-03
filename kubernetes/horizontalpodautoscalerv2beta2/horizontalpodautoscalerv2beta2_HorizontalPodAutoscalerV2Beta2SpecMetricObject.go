package horizontalpodautoscalerv2beta2


type HorizontalPodAutoscalerV2Beta2SpecMetricObject struct {
	// described_object block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler_v2beta2#described_object HorizontalPodAutoscalerV2Beta2#described_object}
	DescribedObject *HorizontalPodAutoscalerV2Beta2SpecMetricObjectDescribedObject `field:"required" json:"describedObject" yaml:"describedObject"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler_v2beta2#metric HorizontalPodAutoscalerV2Beta2#metric}
	Metric *HorizontalPodAutoscalerV2Beta2SpecMetricObjectMetric `field:"required" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler_v2beta2#target HorizontalPodAutoscalerV2Beta2#target}
	Target *HorizontalPodAutoscalerV2Beta2SpecMetricObjectTarget `field:"optional" json:"target" yaml:"target"`
}

