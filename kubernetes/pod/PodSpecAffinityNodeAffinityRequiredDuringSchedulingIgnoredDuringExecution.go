package pod


type PodSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	// node_selector_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod#node_selector_term Pod#node_selector_term}
	NodeSelectorTerm interface{} `field:"optional" json:"nodeSelectorTerm" yaml:"nodeSelectorTerm"`
}

