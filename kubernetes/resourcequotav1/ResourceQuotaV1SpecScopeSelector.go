package resourcequotav1


type ResourceQuotaV1SpecScopeSelector struct {
	// match_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/resource_quota_v1#match_expression ResourceQuotaV1#match_expression}
	MatchExpression interface{} `field:"optional" json:"matchExpression" yaml:"matchExpression"`
}

