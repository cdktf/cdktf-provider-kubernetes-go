// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type StorageClassAllowedTopologies struct {
	// match_label_expressions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class#match_label_expressions StorageClass#match_label_expressions}
	MatchLabelExpressions interface{} `field:"optional" json:"matchLabelExpressions" yaml:"matchLabelExpressions"`
}

