// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecSecurityContextSysctl struct {
	// Name of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job#name CronJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job#value CronJob#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

