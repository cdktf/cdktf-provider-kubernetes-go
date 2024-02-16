// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manifest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManifestConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A Kubernetes manifest describing the desired state of the resource in HCL format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/manifest#manifest Manifest#manifest}
	Manifest *map[string]interface{} `field:"required" json:"manifest" yaml:"manifest"`
	// List of manifest fields whose values can be altered by the API server during 'apply'. Defaults to: ["metadata.annotations", "metadata.labels"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/manifest#computed_fields Manifest#computed_fields}
	ComputedFields *[]*string `field:"optional" json:"computedFields" yaml:"computedFields"`
	// field_manager block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/manifest#field_manager Manifest#field_manager}
	FieldManager *ManifestFieldManager `field:"optional" json:"fieldManager" yaml:"fieldManager"`
	// The resulting resource state, as returned by the API server after applying the desired state from `manifest`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/manifest#object Manifest#object}
	Object *map[string]interface{} `field:"optional" json:"object" yaml:"object"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/manifest#timeouts Manifest#timeouts}
	Timeouts *ManifestTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// wait block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/manifest#wait Manifest#wait}
	Wait *ManifestWait `field:"optional" json:"wait" yaml:"wait"`
	// A map of attribute paths and desired patterns to be matched.
	//
	// After each apply the provider will wait for all attributes listed here to reach a value that matches the desired pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/manifest#wait_for Manifest#wait_for}
	WaitFor *ManifestWaitFor `field:"optional" json:"waitFor" yaml:"waitFor"`
}

