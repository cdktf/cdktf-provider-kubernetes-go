// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecTemplateSpecVolumeProjectedSources struct {
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set_v1#config_map StatefulSetV1#config_map}
	ConfigMap interface{} `field:"optional" json:"configMap" yaml:"configMap"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set_v1#downward_api StatefulSetV1#downward_api}
	DownwardApi *StatefulSetV1SpecTemplateSpecVolumeProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set_v1#secret StatefulSetV1#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// service_account_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set_v1#service_account_token StatefulSetV1#service_account_token}
	ServiceAccountToken *StatefulSetV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

