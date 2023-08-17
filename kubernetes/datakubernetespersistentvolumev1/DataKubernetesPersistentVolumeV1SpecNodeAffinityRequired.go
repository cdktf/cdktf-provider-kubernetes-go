package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecNodeAffinityRequired struct {
	// node_selector_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/data-sources/persistent_volume_v1#node_selector_term DataKubernetesPersistentVolumeV1#node_selector_term}
	NodeSelectorTerm interface{} `field:"required" json:"nodeSelectorTerm" yaml:"nodeSelectorTerm"`
}

