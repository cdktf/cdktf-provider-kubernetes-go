// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type StatefulSetSpecUpdateStrategyRollingUpdate struct {
	// Indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#partition StatefulSet#partition}
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
}

