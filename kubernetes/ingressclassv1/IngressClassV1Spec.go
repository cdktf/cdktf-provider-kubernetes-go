// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ingressclassv1


type IngressClassV1Spec struct {
	// controller refers to the name of the controller that should handle this class.
	//
	// This allows for different "flavors" that are controlled by the same controller. For example, you may have different parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_class_v1#controller IngressClassV1#controller}
	Controller *string `field:"optional" json:"controller" yaml:"controller"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/ingress_class_v1#parameters IngressClassV1#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

