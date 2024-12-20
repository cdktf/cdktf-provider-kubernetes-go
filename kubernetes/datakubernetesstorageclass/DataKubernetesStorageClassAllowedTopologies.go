// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesstorageclass


type DataKubernetesStorageClassAllowedTopologies struct {
	// match_label_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#match_label_expressions DataKubernetesStorageClass#match_label_expressions}
	MatchLabelExpressions interface{} `field:"optional" json:"matchLabelExpressions" yaml:"matchLabelExpressions"`
}

