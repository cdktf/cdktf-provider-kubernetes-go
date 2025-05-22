// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#metadata CronJobV1#metadata}
	Metadata *CronJobV1SpecJobTemplateSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#spec CronJobV1#spec}
	Spec *CronJobV1SpecJobTemplateSpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

