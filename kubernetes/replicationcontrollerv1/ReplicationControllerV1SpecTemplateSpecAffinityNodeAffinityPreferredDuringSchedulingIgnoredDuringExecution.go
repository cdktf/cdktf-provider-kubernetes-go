package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// preference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller_v1#preference ReplicationControllerV1#preference}
	Preference *ReplicationControllerV1SpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	// weight is in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller_v1#weight ReplicationControllerV1#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

