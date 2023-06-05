package statefulsetv1


type StatefulSetV1Spec struct {
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/stateful_set_v1#selector StatefulSetV1#selector}
	Selector *StatefulSetV1SpecSelector `field:"required" json:"selector" yaml:"selector"`
	// The name of the service that governs this StatefulSet.
	//
	// This service must exist before the StatefulSet, and is responsible for the network identity of the set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/stateful_set_v1#service_name StatefulSetV1#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/stateful_set_v1#template StatefulSetV1#template}
	Template *StatefulSetV1SpecTemplate `field:"required" json:"template" yaml:"template"`
	// Controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/stateful_set_v1#pod_management_policy StatefulSetV1#pod_management_policy}
	PodManagementPolicy *string `field:"optional" json:"podManagementPolicy" yaml:"podManagementPolicy"`
	// The desired number of replicas of the given Template, in the sense that they are instantiations of the same Template.
	//
	// Value must be a positive integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/stateful_set_v1#replicas StatefulSetV1#replicas}
	Replicas *string `field:"optional" json:"replicas" yaml:"replicas"`
	// The maximum number of revisions that will be maintained in the StatefulSet's revision history. The default value is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/stateful_set_v1#revision_history_limit StatefulSetV1#revision_history_limit}
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	// update_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/stateful_set_v1#update_strategy StatefulSetV1#update_strategy}
	UpdateStrategy interface{} `field:"optional" json:"updateStrategy" yaml:"updateStrategy"`
	// volume_claim_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/stateful_set_v1#volume_claim_template StatefulSetV1#volume_claim_template}
	VolumeClaimTemplate interface{} `field:"optional" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

