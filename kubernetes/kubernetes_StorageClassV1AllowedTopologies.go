// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type StorageClassV1AllowedTopologies struct {
	// match_label_expressions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#match_label_expressions StorageClassV1#match_label_expressions}
	MatchLabelExpressions interface{} `field:"optional" json:"matchLabelExpressions" yaml:"matchLabelExpressions"`
}

