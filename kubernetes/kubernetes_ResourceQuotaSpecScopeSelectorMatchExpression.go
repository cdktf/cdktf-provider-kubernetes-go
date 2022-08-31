// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ResourceQuotaSpecScopeSelectorMatchExpression struct {
	// Represents a scope's relationship to a set of values.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#operator ResourceQuota#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The name of the scope that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#scope_name ResourceQuota#scope_name}
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
	// A list of scope selector requirements by scope of the resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#values ResourceQuota#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

