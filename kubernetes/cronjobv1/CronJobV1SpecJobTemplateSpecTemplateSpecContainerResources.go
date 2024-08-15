// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecContainerResources struct {
	// Describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job_v1#limits CronJobV1#limits}
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job_v1#requests CronJobV1#requests}
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

