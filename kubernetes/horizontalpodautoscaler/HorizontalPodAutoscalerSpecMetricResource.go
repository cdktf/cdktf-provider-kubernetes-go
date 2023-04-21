package horizontalpodautoscaler


type HorizontalPodAutoscalerSpecMetricResource struct {
	// name is the name of the resource in question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler#name HorizontalPodAutoscaler#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/horizontal_pod_autoscaler#target HorizontalPodAutoscaler#target}
	Target *HorizontalPodAutoscalerSpecMetricResourceTarget `field:"optional" json:"target" yaml:"target"`
}

