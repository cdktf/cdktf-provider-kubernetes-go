package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecMetricPods struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler_v2#metric HorizontalPodAutoscalerV2#metric}
	Metric *HorizontalPodAutoscalerV2SpecMetricPodsMetric `field:"required" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler_v2#target HorizontalPodAutoscalerV2#target}
	Target *HorizontalPodAutoscalerV2SpecMetricPodsTarget `field:"optional" json:"target" yaml:"target"`
}

