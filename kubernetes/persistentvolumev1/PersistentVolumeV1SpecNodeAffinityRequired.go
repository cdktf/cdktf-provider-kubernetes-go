package persistentvolumev1


type PersistentVolumeV1SpecNodeAffinityRequired struct {
	// node_selector_term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#node_selector_term PersistentVolumeV1#node_selector_term}
	NodeSelectorTerm interface{} `field:"required" json:"nodeSelectorTerm" yaml:"nodeSelectorTerm"`
}

