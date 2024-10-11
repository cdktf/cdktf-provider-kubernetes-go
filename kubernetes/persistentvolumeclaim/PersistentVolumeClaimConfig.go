// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumeclaim

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeClaimConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim#metadata PersistentVolumeClaim#metadata}
	Metadata *PersistentVolumeClaimMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim#spec PersistentVolumeClaim#spec}
	Spec *PersistentVolumeClaimSpec `field:"required" json:"spec" yaml:"spec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim#id PersistentVolumeClaim#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim#timeouts PersistentVolumeClaim#timeouts}
	Timeouts *PersistentVolumeClaimTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Whether to wait for the claim to reach `Bound` state (to find volume in which to claim the space).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/persistent_volume_claim#wait_until_bound PersistentVolumeClaim#wait_until_bound}
	WaitUntilBound interface{} `field:"optional" json:"waitUntilBound" yaml:"waitUntilBound"`
}

