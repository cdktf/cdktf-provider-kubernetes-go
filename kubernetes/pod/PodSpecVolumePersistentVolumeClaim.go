// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecVolumePersistentVolumeClaim struct {
	// ClaimName is the name of a PersistentVolumeClaim in the same.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#claim_name Pod#claim_name}
	ClaimName *string `field:"optional" json:"claimName" yaml:"claimName"`
	// Will force the ReadOnly setting in VolumeMounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod#read_only Pod#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

