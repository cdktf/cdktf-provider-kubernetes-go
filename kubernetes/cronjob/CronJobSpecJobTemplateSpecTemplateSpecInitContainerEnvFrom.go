// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecInitContainerEnvFrom struct {
	// config_map_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job#config_map_ref CronJob#config_map_ref}
	ConfigMapRef *CronJobSpecJobTemplateSpecTemplateSpecInitContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job#prefix CronJob#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job#secret_ref CronJob#secret_ref}
	SecretRef *CronJobSpecJobTemplateSpecTemplateSpecInitContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

