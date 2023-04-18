package role


type RoleRule struct {
	// Name of the APIGroup that contains the resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/role#api_groups Role#api_groups}
	ApiGroups *[]*string `field:"required" json:"apiGroups" yaml:"apiGroups"`
	// List of resources that the rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/role#resources Role#resources}
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
	// List of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/role#verbs Role#verbs}
	Verbs *[]*string `field:"required" json:"verbs" yaml:"verbs"`
	// White list of names that the rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/role#resource_names Role#resource_names}
	ResourceNames *[]*string `field:"optional" json:"resourceNames" yaml:"resourceNames"`
}

