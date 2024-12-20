// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manifest


type ManifestWait struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/manifest#condition Manifest#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// A map of paths to fields to wait for a specific field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/manifest#fields Manifest#fields}
	Fields *map[string]*string `field:"optional" json:"fields" yaml:"fields"`
	// Wait for rollout to complete on resources that support `kubectl rollout status`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/manifest#rollout Manifest#rollout}
	Rollout interface{} `field:"optional" json:"rollout" yaml:"rollout"`
}

