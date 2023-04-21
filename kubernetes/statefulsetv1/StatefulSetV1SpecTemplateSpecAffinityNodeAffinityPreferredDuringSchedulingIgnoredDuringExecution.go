package statefulsetv1


type StatefulSetV1SpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// preference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#preference StatefulSetV1#preference}
	Preference *StatefulSetV1SpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference `field:"required" json:"preference" yaml:"preference"`
	// weight is in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#weight StatefulSetV1#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

