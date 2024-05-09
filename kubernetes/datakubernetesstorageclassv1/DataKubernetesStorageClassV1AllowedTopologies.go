// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesstorageclassv1


type DataKubernetesStorageClassV1AllowedTopologies struct {
	// match_label_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/data-sources/storage_class_v1#match_label_expressions DataKubernetesStorageClassV1#match_label_expressions}
	MatchLabelExpressions interface{} `field:"optional" json:"matchLabelExpressions" yaml:"matchLabelExpressions"`
}

