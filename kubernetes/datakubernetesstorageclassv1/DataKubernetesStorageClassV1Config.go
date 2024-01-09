// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetesstorageclassv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataKubernetesStorageClassV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/storage_class_v1#metadata DataKubernetesStorageClassV1#metadata}
	Metadata *DataKubernetesStorageClassV1Metadata `field:"required" json:"metadata" yaml:"metadata"`
	// allowed_topologies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/storage_class_v1#allowed_topologies DataKubernetesStorageClassV1#allowed_topologies}
	AllowedTopologies *DataKubernetesStorageClassV1AllowedTopologies `field:"optional" json:"allowedTopologies" yaml:"allowedTopologies"`
	// Indicates whether the storage class allow volume expand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/storage_class_v1#allow_volume_expansion DataKubernetesStorageClassV1#allow_volume_expansion}
	AllowVolumeExpansion interface{} `field:"optional" json:"allowVolumeExpansion" yaml:"allowVolumeExpansion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/storage_class_v1#id DataKubernetesStorageClassV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Persistent Volumes that are dynamically created by a storage class will have the mount options specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/storage_class_v1#mount_options DataKubernetesStorageClassV1#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The parameters for the provisioner that should create volumes of this storage class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/storage_class_v1#parameters DataKubernetesStorageClassV1#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Indicates the type of the reclaim policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/storage_class_v1#reclaim_policy DataKubernetesStorageClassV1#reclaim_policy}
	ReclaimPolicy *string `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// Indicates when volume binding and dynamic provisioning should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/data-sources/storage_class_v1#volume_binding_mode DataKubernetesStorageClassV1#volume_binding_mode}
	VolumeBindingMode *string `field:"optional" json:"volumeBindingMode" yaml:"volumeBindingMode"`
}

