// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datakubernetespersistentvolumev1


type DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbd struct {
	// A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#ceph_monitors DataKubernetesPersistentVolumeV1#ceph_monitors}
	CephMonitors *[]*string `field:"required" json:"cephMonitors" yaml:"cephMonitors"`
	// The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#rbd_image DataKubernetesPersistentVolumeV1#rbd_image}
	RbdImage *string `field:"required" json:"rbdImage" yaml:"rbdImage"`
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#fs_type DataKubernetesPersistentVolumeV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#keyring DataKubernetesPersistentVolumeV1#keyring}
	Keyring *string `field:"optional" json:"keyring" yaml:"keyring"`
	// The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#rados_user DataKubernetesPersistentVolumeV1#rados_user}
	RadosUser *string `field:"optional" json:"radosUser" yaml:"radosUser"`
	// The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#rbd_pool DataKubernetesPersistentVolumeV1#rbd_pool}
	RbdPool *string `field:"optional" json:"rbdPool" yaml:"rbdPool"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#read_only DataKubernetesPersistentVolumeV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/data-sources/persistent_volume_v1#secret_ref DataKubernetesPersistentVolumeV1#secret_ref}
	SecretRef *DataKubernetesPersistentVolumeV1SpecPersistentVolumeSourceRbdSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

