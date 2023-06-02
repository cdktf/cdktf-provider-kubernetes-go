package nodetaint


type NodeTaintMetadata struct {
	// The name of the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/node_taint#name NodeTaint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

