package daemonset


type DaemonsetSpecTemplateSpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/daemonset#condition_type Daemonset#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

