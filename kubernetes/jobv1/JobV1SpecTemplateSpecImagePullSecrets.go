// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecTemplateSpecImagePullSecrets struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/job_v1#name JobV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

