// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecAffinity struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job_v1#node_affinity CronJobV1#node_affinity}
	NodeAffinity *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// pod_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job_v1#pod_affinity CronJobV1#pod_affinity}
	PodAffinity *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_anti_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/cron_job_v1#pod_anti_affinity CronJobV1#pod_anti_affinity}
	PodAntiAffinity *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

