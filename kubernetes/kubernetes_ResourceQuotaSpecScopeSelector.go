// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ResourceQuotaSpecScopeSelector struct {
	// match_expression block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#match_expression ResourceQuota#match_expression}
	MatchExpression interface{} `field:"optional" json:"matchExpression" yaml:"matchExpression"`
}

