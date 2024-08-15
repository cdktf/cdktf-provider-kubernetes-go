// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job#delete CronJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

