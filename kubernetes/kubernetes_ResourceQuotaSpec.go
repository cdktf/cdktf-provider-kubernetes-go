// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ResourceQuotaSpec struct {
	// The set of desired hard limits for each named resource. More info: http://releases.k8s.io/HEAD/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#hard ResourceQuota#hard}
	Hard *map[string]*string `field:"optional" json:"hard" yaml:"hard"`
	// A collection of filters that must match each object tracked by a quota.
	//
	// If not specified, the quota matches all objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#scopes ResourceQuota#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// scope_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#scope_selector ResourceQuota#scope_selector}
	ScopeSelector *ResourceQuotaSpecScopeSelector `field:"optional" json:"scopeSelector" yaml:"scopeSelector"`
}
