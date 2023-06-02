package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecMetricExternal struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/horizontal_pod_autoscaler#metric HorizontalPodAutoscaler#metric}
	Metric *HorizontalPodAutoscalerSpecMetricExternalMetric `field:"required" json:"metric" yaml:"metric"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/horizontal_pod_autoscaler#target HorizontalPodAutoscaler#target}
	Target *HorizontalPodAutoscalerSpecMetricExternalTarget `field:"optional" json:"target" yaml:"target"`
}

