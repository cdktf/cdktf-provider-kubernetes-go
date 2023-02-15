package datakubernetesstorageclassv1


type DataKubernetesStorageClassV1AllowedTopologies struct {
	// match_label_expressions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/storage_class_v1#match_label_expressions DataKubernetesStorageClassV1#match_label_expressions}
	MatchLabelExpressions interface{} `field:"optional" json:"matchLabelExpressions" yaml:"matchLabelExpressions"`
}

