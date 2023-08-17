package podv1


type PodV1SpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference struct {
	// match_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod_v1#match_expressions PodV1#match_expressions}
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
}

