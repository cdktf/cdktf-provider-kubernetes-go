package replicationcontroller


type ReplicationControllerSpecTemplateSpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/replication_controller#condition_type ReplicationController#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}
