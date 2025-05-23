// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type KubernetesProviderExperiments struct {
	// Enable the `kubernetes_manifest` resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs#manifest_resource KubernetesProvider#manifest_resource}
	ManifestResource interface{} `field:"optional" json:"manifestResource" yaml:"manifestResource"`
}

