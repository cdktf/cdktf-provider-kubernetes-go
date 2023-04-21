package job


type JobSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference struct {
	// match_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#match_expressions Job#match_expressions}
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
}

