// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecVolumeProjectedSources struct {
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/pod_v1#config_map PodV1#config_map}
	ConfigMap interface{} `field:"optional" json:"configMap" yaml:"configMap"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/pod_v1#downward_api PodV1#downward_api}
	DownwardApi *PodV1SpecVolumeProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/pod_v1#secret PodV1#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// service_account_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/pod_v1#service_account_token PodV1#service_account_token}
	ServiceAccountToken *PodV1SpecVolumeProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

