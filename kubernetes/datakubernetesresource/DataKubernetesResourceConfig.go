// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesresource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesResourceConfig struct {
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
	// The resource apiVersion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/data-sources/resource#api_version DataKubernetesResource#api_version}
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// The resource kind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/data-sources/resource#kind DataKubernetesResource#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/data-sources/resource#metadata DataKubernetesResource#metadata}
	Metadata *DataKubernetesResourceMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// The response from the API server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/data-sources/resource#object DataKubernetesResource#object}
	Object *map[string]interface{} `field:"optional" json:"object" yaml:"object"`
}

