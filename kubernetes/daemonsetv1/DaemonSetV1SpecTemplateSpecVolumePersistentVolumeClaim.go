// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1


type DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaim struct {
	// ClaimName is the name of a PersistentVolumeClaim in the same.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/daemon_set_v1#claim_name DaemonSetV1#claim_name}
	ClaimName *string `field:"optional" json:"claimName" yaml:"claimName"`
	// Will force the ReadOnly setting in VolumeMounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/daemon_set_v1#read_only DaemonSetV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

