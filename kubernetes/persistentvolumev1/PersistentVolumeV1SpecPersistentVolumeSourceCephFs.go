// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumev1


type PersistentVolumeV1SpecPersistentVolumeSourceCephFs struct {
	// Monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#monitors PersistentVolumeV1#monitors}
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	// Used as the mounted root, rather than the full Ceph tree, default is /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#path PersistentVolumeV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to `false` (read/write). More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#read_only PersistentVolumeV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The path to key ring for User, default is `/etc/ceph/user.secret`. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#secret_file PersistentVolumeV1#secret_file}
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#secret_ref PersistentVolumeV1#secret_ref}
	SecretRef *PersistentVolumeV1SpecPersistentVolumeSourceCephFsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// User is the rados user name, default is admin. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/persistent_volume_v1#user PersistentVolumeV1#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

