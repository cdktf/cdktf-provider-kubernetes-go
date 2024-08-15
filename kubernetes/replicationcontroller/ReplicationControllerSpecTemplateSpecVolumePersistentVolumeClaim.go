// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller


type ReplicationControllerSpecTemplateSpecVolumePersistentVolumeClaim struct {
	// ClaimName is the name of a PersistentVolumeClaim in the same.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/replication_controller#claim_name ReplicationController#claim_name}
	ClaimName *string `field:"optional" json:"claimName" yaml:"claimName"`
	// Will force the ReadOnly setting in VolumeMounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/replication_controller#read_only ReplicationController#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

