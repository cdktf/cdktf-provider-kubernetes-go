package datakubernetesnodes


type DataKubernetesNodesMetadata struct {
	// Select nodes with these labels. More info: http://kubernetes.io/docs/user-guide/labels.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/nodes#labels DataKubernetesNodes#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
}

