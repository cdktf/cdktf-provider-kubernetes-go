// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentv1


type DeploymentV1SpecTemplateSpecVolumeProjectedSources struct {
	// config_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/deployment_v1#config_map DeploymentV1#config_map}
	ConfigMap interface{} `field:"optional" json:"configMap" yaml:"configMap"`
	// downward_api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/deployment_v1#downward_api DeploymentV1#downward_api}
	DownwardApi *DeploymentV1SpecTemplateSpecVolumeProjectedSourcesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/deployment_v1#secret DeploymentV1#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// service_account_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/deployment_v1#service_account_token DeploymentV1#service_account_token}
	ServiceAccountToken *DeploymentV1SpecTemplateSpecVolumeProjectedSourcesServiceAccountToken `field:"optional" json:"serviceAccountToken" yaml:"serviceAccountToken"`
}

