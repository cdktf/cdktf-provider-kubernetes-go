package resourcequota


type ResourceQuotaSpecScopeSelectorMatchExpression struct {
	// Represents a scope's relationship to a set of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/resource_quota#operator ResourceQuota#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The name of the scope that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/resource_quota#scope_name ResourceQuota#scope_name}
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
	// A list of scope selector requirements by scope of the resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/resource_quota#values ResourceQuota#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

