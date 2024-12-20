// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesstorageclass

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesStorageClassConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#metadata DataKubernetesStorageClass#metadata}
	Metadata *DataKubernetesStorageClassMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// allowed_topologies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#allowed_topologies DataKubernetesStorageClass#allowed_topologies}
	AllowedTopologies *DataKubernetesStorageClassAllowedTopologies `field:"optional" json:"allowedTopologies" yaml:"allowedTopologies"`
	// Indicates whether the storage class allow volume expand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#allow_volume_expansion DataKubernetesStorageClass#allow_volume_expansion}
	AllowVolumeExpansion interface{} `field:"optional" json:"allowVolumeExpansion" yaml:"allowVolumeExpansion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#id DataKubernetesStorageClass#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Persistent Volumes that are dynamically created by a storage class will have the mount options specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#mount_options DataKubernetesStorageClass#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The parameters for the provisioner that should create volumes of this storage class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#parameters DataKubernetesStorageClass#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Indicates the type of the reclaim policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#reclaim_policy DataKubernetesStorageClass#reclaim_policy}
	ReclaimPolicy *string `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// Indicates when volume binding and dynamic provisioning should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/data-sources/storage_class#volume_binding_mode DataKubernetesStorageClass#volume_binding_mode}
	VolumeBindingMode *string `field:"optional" json:"volumeBindingMode" yaml:"volumeBindingMode"`
}

