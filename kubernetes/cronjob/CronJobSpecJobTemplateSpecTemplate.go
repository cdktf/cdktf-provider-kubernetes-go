// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/cron_job#metadata CronJob#metadata}
	Metadata *CronJobSpecJobTemplateSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/cron_job#spec CronJob#spec}
	Spec *CronJobSpecJobTemplateSpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

