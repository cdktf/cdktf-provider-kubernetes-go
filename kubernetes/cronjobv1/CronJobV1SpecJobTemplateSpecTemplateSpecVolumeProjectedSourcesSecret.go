// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecret struct {
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job_v1#items CronJobV1#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job_v1#name CronJobV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional: Specify whether the Secret or it's keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job_v1#optional CronJobV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

