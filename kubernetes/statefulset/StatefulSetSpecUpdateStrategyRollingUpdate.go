package statefulset


type StatefulSetSpecUpdateStrategyRollingUpdate struct {
	// Indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set#partition StatefulSet#partition}
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
}

