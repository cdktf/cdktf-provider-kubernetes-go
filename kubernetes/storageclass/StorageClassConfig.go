// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageclass

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageClassConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#metadata StorageClass#metadata}
	Metadata *StorageClassMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Indicates the type of the provisioner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#storage_provisioner StorageClass#storage_provisioner}
	StorageProvisioner *string `field:"required" json:"storageProvisioner" yaml:"storageProvisioner"`
	// allowed_topologies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#allowed_topologies StorageClass#allowed_topologies}
	AllowedTopologies *StorageClassAllowedTopologies `field:"optional" json:"allowedTopologies" yaml:"allowedTopologies"`
	// Indicates whether the storage class allow volume expand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#allow_volume_expansion StorageClass#allow_volume_expansion}
	AllowVolumeExpansion interface{} `field:"optional" json:"allowVolumeExpansion" yaml:"allowVolumeExpansion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#id StorageClass#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Persistent Volumes that are dynamically created by a storage class will have the mount options specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#mount_options StorageClass#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The parameters for the provisioner that should create volumes of this storage class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#parameters StorageClass#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Indicates the type of the reclaim policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#reclaim_policy StorageClass#reclaim_policy}
	ReclaimPolicy *string `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// Indicates when volume binding and dynamic provisioning should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/storage_class#volume_binding_mode StorageClass#volume_binding_mode}
	VolumeBindingMode *string `field:"optional" json:"volumeBindingMode" yaml:"volumeBindingMode"`
}

