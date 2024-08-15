// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namespacev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NamespaceV1Config struct {
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
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/namespace_v1#metadata NamespaceV1#metadata}
	Metadata *NamespaceV1Metadata `field:"required" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/namespace_v1#id NamespaceV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/namespace_v1#timeouts NamespaceV1#timeouts}
	Timeouts *NamespaceV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Terraform will wait for the default service account to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/namespace_v1#wait_for_default_service_account NamespaceV1#wait_for_default_service_account}
	WaitForDefaultServiceAccount interface{} `field:"optional" json:"waitForDefaultServiceAccount" yaml:"waitForDefaultServiceAccount"`
}

