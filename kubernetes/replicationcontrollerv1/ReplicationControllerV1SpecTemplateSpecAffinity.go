package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecAffinity struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#node_affinity ReplicationControllerV1#node_affinity}
	NodeAffinity *ReplicationControllerV1SpecTemplateSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// pod_affinity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#pod_affinity ReplicationControllerV1#pod_affinity}
	PodAffinity *ReplicationControllerV1SpecTemplateSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_anti_affinity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#pod_anti_affinity ReplicationControllerV1#pod_anti_affinity}
	PodAntiAffinity *ReplicationControllerV1SpecTemplateSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

