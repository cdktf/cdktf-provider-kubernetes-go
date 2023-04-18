package clusterrole


type ClusterRoleAggregationRule struct {
	// cluster_role_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/cluster_role#cluster_role_selectors ClusterRole#cluster_role_selectors}
	ClusterRoleSelectors interface{} `field:"optional" json:"clusterRoleSelectors" yaml:"clusterRoleSelectors"`
}

