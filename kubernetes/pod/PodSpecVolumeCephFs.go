// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecVolumeCephFs struct {
	// Monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#monitors Pod#monitors}
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	// Used as the mounted root, rather than the full Ceph tree, default is /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#path Pod#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to `false` (read/write). More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#read_only Pod#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The path to key ring for User, default is `/etc/ceph/user.secret`. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#secret_file Pod#secret_file}
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#secret_ref Pod#secret_ref}
	SecretRef *PodSpecVolumeCephFsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// User is the rados user name, default is admin. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#user Pod#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

