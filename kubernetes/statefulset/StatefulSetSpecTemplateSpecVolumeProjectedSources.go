// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecVolumeProjectedSources struct {
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#config_map StatefulSet#config_map}
	ConfigMap interface{} `field:"optional" json:"configMap" yaml:"configMap"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#downward_api StatefulSet#downward_api}
	DownwardApi *StatefulSetSpecTemplateSpecVolumeProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#secret StatefulSet#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// service_account_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#service_account_token StatefulSet#service_account_token}
	ServiceAccountToken *StatefulSetSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

