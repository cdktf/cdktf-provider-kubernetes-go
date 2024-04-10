// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaultserviceaccountv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DefaultServiceAccountV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/default_service_account_v1#metadata DefaultServiceAccountV1#metadata}
	Metadata *DefaultServiceAccountV1Metadata `field:"required" json:"metadata" yaml:"metadata"`
	// Enable automatic mounting of the service account token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/default_service_account_v1#automount_service_account_token DefaultServiceAccountV1#automount_service_account_token}
	AutomountServiceAccountToken interface{} `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/default_service_account_v1#id DefaultServiceAccountV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image_pull_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/default_service_account_v1#image_pull_secret DefaultServiceAccountV1#image_pull_secret}
	ImagePullSecret interface{} `field:"optional" json:"imagePullSecret" yaml:"imagePullSecret"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/default_service_account_v1#secret DefaultServiceAccountV1#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/default_service_account_v1#timeouts DefaultServiceAccountV1#timeouts}
	Timeouts *DefaultServiceAccountV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

