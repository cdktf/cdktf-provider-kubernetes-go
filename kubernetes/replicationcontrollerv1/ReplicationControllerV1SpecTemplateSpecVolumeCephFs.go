// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecVolumeCephFs struct {
	// Monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/replication_controller_v1#monitors ReplicationControllerV1#monitors}
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	// Used as the mounted root, rather than the full Ceph tree, default is /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/replication_controller_v1#path ReplicationControllerV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to `false` (read/write). More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/replication_controller_v1#read_only ReplicationControllerV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The path to key ring for User, default is `/etc/ceph/user.secret`. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/replication_controller_v1#secret_file ReplicationControllerV1#secret_file}
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/replication_controller_v1#secret_ref ReplicationControllerV1#secret_ref}
	SecretRef *ReplicationControllerV1SpecTemplateSpecVolumeCephFsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// User is the rados user name, default is admin. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs/resources/replication_controller_v1#user ReplicationControllerV1#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

