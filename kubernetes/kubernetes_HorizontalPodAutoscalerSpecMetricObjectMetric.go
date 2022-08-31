// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type HorizontalPodAutoscalerSpecMetricObjectMetric struct {
	// name is the name of the given metric.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler#name HorizontalPodAutoscaler#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler#selector HorizontalPodAutoscaler#selector}
	Selector interface{} `field:"optional" json:"selector" yaml:"selector"`
}

