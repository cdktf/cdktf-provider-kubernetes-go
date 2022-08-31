// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#condition_type Pod#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

