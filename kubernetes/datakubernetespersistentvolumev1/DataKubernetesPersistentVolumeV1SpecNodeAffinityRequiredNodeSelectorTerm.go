// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecNodeAffinityRequiredNodeSelectorTerm struct {
	// match_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/data-sources/persistent_volume_v1#match_expressions DataKubernetesPersistentVolumeV1#match_expressions}
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// match_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/data-sources/persistent_volume_v1#match_fields DataKubernetesPersistentVolumeV1#match_fields}
	MatchFields interface{} `field:"optional" json:"matchFields" yaml:"matchFields"`
}

