package nodetaint


type NodeTaintMetadata struct {
	// The name of the node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/node_taint#name NodeTaint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

