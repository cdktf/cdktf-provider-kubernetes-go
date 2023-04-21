package daemonsetv1


type DaemonSetV1SpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	// node_selector_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemon_set_v1#node_selector_term DaemonSetV1#node_selector_term}
	NodeSelectorTerm interface{} `field:"optional" json:"nodeSelectorTerm" yaml:"nodeSelectorTerm"`
}

