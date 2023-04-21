package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	// node_selector_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller_v1#node_selector_term ReplicationControllerV1#node_selector_term}
	NodeSelectorTerm interface{} `field:"optional" json:"nodeSelectorTerm" yaml:"nodeSelectorTerm"`
}

