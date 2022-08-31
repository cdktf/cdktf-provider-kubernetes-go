// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type DataKubernetesStorageClassAllowedTopologies struct {
	// match_label_expressions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/storage_class#match_label_expressions DataKubernetesStorageClass#match_label_expressions}
	MatchLabelExpressions interface{} `field:"optional" json:"matchLabelExpressions" yaml:"matchLabelExpressions"`
}

