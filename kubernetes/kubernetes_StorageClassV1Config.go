// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageClassV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#metadata StorageClassV1#metadata}
	Metadata *StorageClassV1Metadata `field:"required" json:"metadata" yaml:"metadata"`
	// Indicates the type of the provisioner.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#storage_provisioner StorageClassV1#storage_provisioner}
	StorageProvisioner *string `field:"required" json:"storageProvisioner" yaml:"storageProvisioner"`
	// allowed_topologies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#allowed_topologies StorageClassV1#allowed_topologies}
	AllowedTopologies *StorageClassV1AllowedTopologies `field:"optional" json:"allowedTopologies" yaml:"allowedTopologies"`
	// Indicates whether the storage class allow volume expand.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#allow_volume_expansion StorageClassV1#allow_volume_expansion}
	AllowVolumeExpansion interface{} `field:"optional" json:"allowVolumeExpansion" yaml:"allowVolumeExpansion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#id StorageClassV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Persistent Volumes that are dynamically created by a storage class will have the mount options specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#mount_options StorageClassV1#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The parameters for the provisioner that should create volumes of this storage class.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#parameters StorageClassV1#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Indicates the type of the reclaim policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#reclaim_policy StorageClassV1#reclaim_policy}
	ReclaimPolicy *string `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// Indicates when volume binding and dynamic provisioning should occur.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class_v1#volume_binding_mode StorageClassV1#volume_binding_mode}
	VolumeBindingMode *string `field:"optional" json:"volumeBindingMode" yaml:"volumeBindingMode"`
}

