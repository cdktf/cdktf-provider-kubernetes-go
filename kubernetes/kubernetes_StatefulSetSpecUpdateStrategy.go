// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type StatefulSetSpecUpdateStrategy struct {
	// rolling_update block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#rolling_update StatefulSet#rolling_update}
	RollingUpdate interface{} `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Indicates the type of the StatefulSet update strategy. Default is RollingUpdate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#type StatefulSet#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

