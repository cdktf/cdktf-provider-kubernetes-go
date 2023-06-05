package nodetaint


type NodeTaintTaint struct {
	// The taint effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/node_taint#effect NodeTaint#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// The taint key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/node_taint#key NodeTaint#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The taint value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/node_taint#value NodeTaint#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

