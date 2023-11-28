// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller


type ReplicationControllerSpecTemplateSpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller#name ReplicationController#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller#value ReplicationController#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

