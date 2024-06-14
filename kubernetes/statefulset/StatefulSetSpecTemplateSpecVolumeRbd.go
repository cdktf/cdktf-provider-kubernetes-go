// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecVolumeRbd struct {
	// A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#ceph_monitors StatefulSet#ceph_monitors}
	CephMonitors *[]*string `field:"required" json:"cephMonitors" yaml:"cephMonitors"`
	// The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#rbd_image StatefulSet#rbd_image}
	RbdImage *string `field:"required" json:"rbdImage" yaml:"rbdImage"`
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#fs_type StatefulSet#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#keyring StatefulSet#keyring}
	Keyring *string `field:"optional" json:"keyring" yaml:"keyring"`
	// The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#rados_user StatefulSet#rados_user}
	RadosUser *string `field:"optional" json:"radosUser" yaml:"radosUser"`
	// The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#rbd_pool StatefulSet#rbd_pool}
	RbdPool *string `field:"optional" json:"rbdPool" yaml:"rbdPool"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#read_only StatefulSet#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/stateful_set#secret_ref StatefulSet#secret_ref}
	SecretRef *StatefulSetSpecTemplateSpecVolumeRbdSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

