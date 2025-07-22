// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manifest


type ManifestWaitFor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/manifest#fields Manifest#fields}.
	Fields *map[string]*string `field:"optional" json:"fields" yaml:"fields"`
}

