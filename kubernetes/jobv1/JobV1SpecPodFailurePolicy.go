// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecPodFailurePolicy struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/job_v1#rule JobV1#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

