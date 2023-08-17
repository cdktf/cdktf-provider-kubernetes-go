package resourcequota


type ResourceQuotaSpecScopeSelector struct {
	// match_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/resource_quota#match_expression ResourceQuota#match_expression}
	MatchExpression interface{} `field:"optional" json:"matchExpression" yaml:"matchExpression"`
}

