// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecTemplateSpecAffinity struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job_v1#node_affinity JobV1#node_affinity}
	NodeAffinity *JobV1SpecTemplateSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// pod_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job_v1#pod_affinity JobV1#pod_affinity}
	PodAffinity *JobV1SpecTemplateSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_anti_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job_v1#pod_anti_affinity JobV1#pod_anti_affinity}
	PodAntiAffinity *JobV1SpecTemplateSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

