package jobv1


type JobV1SpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	// node_selector_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/job_v1#node_selector_term JobV1#node_selector_term}
	NodeSelectorTerm interface{} `field:"optional" json:"nodeSelectorTerm" yaml:"nodeSelectorTerm"`
}

