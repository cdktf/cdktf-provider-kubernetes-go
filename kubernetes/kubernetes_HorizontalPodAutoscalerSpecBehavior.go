// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type HorizontalPodAutoscalerSpecBehavior struct {
	// scale_down block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler#scale_down HorizontalPodAutoscaler#scale_down}
	ScaleDown interface{} `field:"optional" json:"scaleDown" yaml:"scaleDown"`
	// scale_up block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler#scale_up HorizontalPodAutoscaler#scale_up}
	ScaleUp interface{} `field:"optional" json:"scaleUp" yaml:"scaleUp"`
}

