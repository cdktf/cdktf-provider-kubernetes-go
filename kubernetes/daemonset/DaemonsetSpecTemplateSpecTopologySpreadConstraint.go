package daemonset


type DaemonsetSpecTemplateSpecTopologySpreadConstraint struct {
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#label_selector Daemonset#label_selector}
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// describes the degree to which pods may be unevenly distributed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#max_skew Daemonset#max_skew}
	MaxSkew *float64 `field:"optional" json:"maxSkew" yaml:"maxSkew"`
	// the key of node labels.
	//
	// Nodes that have a label with this key and identical values are considered to be in the same topology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#topology_key Daemonset#topology_key}
	TopologyKey *string `field:"optional" json:"topologyKey" yaml:"topologyKey"`
	// indicates how to deal with a pod if it doesn't satisfy the spread constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#when_unsatisfiable Daemonset#when_unsatisfiable}
	WhenUnsatisfiable *string `field:"optional" json:"whenUnsatisfiable" yaml:"whenUnsatisfiable"`
}

