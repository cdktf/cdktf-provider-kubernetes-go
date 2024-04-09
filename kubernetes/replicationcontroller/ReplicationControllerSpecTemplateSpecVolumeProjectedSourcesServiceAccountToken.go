// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller


type ReplicationControllerSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken struct {
	// Path specifies a relative path to the mount point of the projected volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/replication_controller#path ReplicationController#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Audience is the intended audience of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/replication_controller#audience ReplicationController#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// ExpirationSeconds is the expected duration of validity of the service account token.
	//
	// It defaults to 1 hour and must be at least 10 minutes (600 seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/replication_controller#expiration_seconds ReplicationController#expiration_seconds}
	ExpirationSeconds *float64 `field:"optional" json:"expirationSeconds" yaml:"expirationSeconds"`
}

