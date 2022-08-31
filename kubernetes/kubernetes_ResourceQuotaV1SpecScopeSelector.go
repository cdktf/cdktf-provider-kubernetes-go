// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ResourceQuotaV1SpecScopeSelector struct {
	// match_expression block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota_v1#match_expression ResourceQuotaV1#match_expression}
	MatchExpression interface{} `field:"optional" json:"matchExpression" yaml:"matchExpression"`
}

