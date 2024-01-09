// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecInitContainerEnvFromConfigMapRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cron_job#name CronJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the ConfigMap must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cron_job#optional CronJob#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

