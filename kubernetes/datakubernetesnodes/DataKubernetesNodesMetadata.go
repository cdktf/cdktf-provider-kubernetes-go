package datakubernetesnodes


type DataKubernetesNodesMetadata struct {
	// Select nodes with these labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/data-sources/nodes#labels DataKubernetesNodes#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
}

