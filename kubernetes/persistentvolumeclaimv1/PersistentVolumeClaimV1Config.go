// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumeclaimv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeClaimV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/persistent_volume_claim_v1#metadata PersistentVolumeClaimV1#metadata}
	Metadata *PersistentVolumeClaimV1Metadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/persistent_volume_claim_v1#spec PersistentVolumeClaimV1#spec}
	Spec *PersistentVolumeClaimV1Spec `field:"required" json:"spec" yaml:"spec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/persistent_volume_claim_v1#id PersistentVolumeClaimV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/persistent_volume_claim_v1#timeouts PersistentVolumeClaimV1#timeouts}
	Timeouts *PersistentVolumeClaimV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Whether to wait for the claim to reach `Bound` state (to find volume in which to claim the space).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/persistent_volume_claim_v1#wait_until_bound PersistentVolumeClaimV1#wait_until_bound}
	WaitUntilBound interface{} `field:"optional" json:"waitUntilBound" yaml:"waitUntilBound"`
}

