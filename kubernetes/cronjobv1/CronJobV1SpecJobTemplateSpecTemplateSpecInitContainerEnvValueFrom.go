// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerEnvValueFrom struct {
	// config_map_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/cron_job_v1#config_map_key_ref CronJobV1#config_map_key_ref}
	ConfigMapKeyRef *CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/cron_job_v1#field_ref CronJobV1#field_ref}
	FieldRef *CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/cron_job_v1#resource_field_ref CronJobV1#resource_field_ref}
	ResourceFieldRef *CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/cron_job_v1#secret_key_ref CronJobV1#secret_key_ref}
	SecretKeyRef *CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

