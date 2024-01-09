// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1


type DaemonSetV1SpecTemplateSpecVolumeProjectedSources struct {
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemon_set_v1#config_map DaemonSetV1#config_map}
	ConfigMap interface{} `field:"optional" json:"configMap" yaml:"configMap"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemon_set_v1#downward_api DaemonSetV1#downward_api}
	DownwardApi *DaemonSetV1SpecTemplateSpecVolumeProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemon_set_v1#secret DaemonSetV1#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// service_account_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemon_set_v1#service_account_token DaemonSetV1#service_account_token}
	ServiceAccountToken *DaemonSetV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

