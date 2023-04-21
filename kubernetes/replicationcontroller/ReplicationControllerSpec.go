package replicationcontroller


type ReplicationControllerSpec struct {
	// A label query over pods that should match the Replicas count.
	//
	// If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller#selector ReplicationController#selector}
	Selector *map[string]*string `field:"required" json:"selector" yaml:"selector"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller#template ReplicationController#template}
	Template *ReplicationControllerSpecTemplate `field:"required" json:"template" yaml:"template"`
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	//
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller#min_ready_seconds ReplicationController#min_ready_seconds}
	MinReadySeconds *float64 `field:"optional" json:"minReadySeconds" yaml:"minReadySeconds"`
	// The number of desired replicas. Defaults to 1. More info: http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller#replicas ReplicationController#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

