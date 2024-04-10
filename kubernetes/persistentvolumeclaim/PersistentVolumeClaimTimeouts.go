// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumeclaim


type PersistentVolumeClaimTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/persistent_volume_claim#create PersistentVolumeClaim#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

