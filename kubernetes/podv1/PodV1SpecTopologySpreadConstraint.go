package podv1


type PodV1SpecTopologySpreadConstraint struct {
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#label_selector PodV1#label_selector}
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// describes the degree to which pods may be unevenly distributed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#max_skew PodV1#max_skew}
	MaxSkew *float64 `field:"optional" json:"maxSkew" yaml:"maxSkew"`
	// the key of node labels.
	//
	// Nodes that have a label with this key and identical values are considered to be in the same topology.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#topology_key PodV1#topology_key}
	TopologyKey *string `field:"optional" json:"topologyKey" yaml:"topologyKey"`
	// indicates how to deal with a pod if it doesn't satisfy the spread constraint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#when_unsatisfiable PodV1#when_unsatisfiable}
	WhenUnsatisfiable *string `field:"optional" json:"whenUnsatisfiable" yaml:"whenUnsatisfiable"`
}

