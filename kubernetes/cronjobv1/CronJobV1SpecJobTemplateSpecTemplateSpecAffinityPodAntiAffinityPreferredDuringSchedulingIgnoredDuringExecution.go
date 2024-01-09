// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// pod_affinity_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cron_job_v1#pod_affinity_term CronJobV1#pod_affinity_term}
	PodAffinityTerm *CronJobV1SpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cron_job_v1#weight CronJobV1#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

