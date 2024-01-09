// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labels

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LabelsConfig struct {
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
	// The apiVersion of the resource to label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/labels#api_version Labels#api_version}
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// The kind of the resource to label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/labels#kind Labels#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// A map of labels to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/labels#labels Labels#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/labels#metadata Labels#metadata}
	Metadata *LabelsMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Set the name of the field manager for the specified labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/labels#field_manager Labels#field_manager}
	FieldManager *string `field:"optional" json:"fieldManager" yaml:"fieldManager"`
	// Force overwriting labels that were created or edited outside of Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/labels#force Labels#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/labels#id Labels#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

