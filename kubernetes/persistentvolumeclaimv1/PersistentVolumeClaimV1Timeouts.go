// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package persistentvolumeclaimv1


type PersistentVolumeClaimV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/persistent_volume_claim_v1#create PersistentVolumeClaimV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

