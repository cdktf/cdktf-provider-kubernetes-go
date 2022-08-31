// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type StatefulSetSpec struct {
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#selector StatefulSet#selector}
	Selector *StatefulSetSpecSelector `field:"required" json:"selector" yaml:"selector"`
	// The name of the service that governs this StatefulSet.
	//
	// This service must exist before the StatefulSet, and is responsible for the network identity of the set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#service_name StatefulSet#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#template StatefulSet#template}
	Template *StatefulSetSpecTemplate `field:"required" json:"template" yaml:"template"`
	// Controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#pod_management_policy StatefulSet#pod_management_policy}
	PodManagementPolicy *string `field:"optional" json:"podManagementPolicy" yaml:"podManagementPolicy"`
	// The desired number of replicas of the given Template, in the sense that they are instantiations of the same Template.
	//
	// Value must be a positive integer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#replicas StatefulSet#replicas}
	Replicas *string `field:"optional" json:"replicas" yaml:"replicas"`
	// The maximum number of revisions that will be maintained in the StatefulSet's revision history. The default value is 10.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#revision_history_limit StatefulSet#revision_history_limit}
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	// update_strategy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#update_strategy StatefulSet#update_strategy}
	UpdateStrategy interface{} `field:"optional" json:"updateStrategy" yaml:"updateStrategy"`
	// volume_claim_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#volume_claim_template StatefulSet#volume_claim_template}
	VolumeClaimTemplate interface{} `field:"optional" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

