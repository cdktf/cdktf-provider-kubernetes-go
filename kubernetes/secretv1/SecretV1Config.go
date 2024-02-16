// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/secret_v1#metadata SecretV1#metadata}
	Metadata *SecretV1Metadata `field:"required" json:"metadata" yaml:"metadata"`
	// A map of the secret data in base64 encoding. Use this for binary data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/secret_v1#binary_data SecretV1#binary_data}
	BinaryData *map[string]*string `field:"optional" json:"binaryData" yaml:"binaryData"`
	// A map of the secret data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/secret_v1#data SecretV1#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/secret_v1#id SecretV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/secret_v1#immutable SecretV1#immutable}
	Immutable interface{} `field:"optional" json:"immutable" yaml:"immutable"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/secret_v1#timeouts SecretV1#timeouts}
	Timeouts *SecretV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Type of secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/secret_v1#type SecretV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Terraform will wait for the service account token to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/secret_v1#wait_for_service_account_token SecretV1#wait_for_service_account_token}
	WaitForServiceAccountToken interface{} `field:"optional" json:"waitForServiceAccountToken" yaml:"waitForServiceAccountToken"`
}

