// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceSpecSessionAffinityConfigClientIp struct {
	// Specifies the seconds of `ClientIP` type session sticky time.
	//
	// The value must be > 0 and <= 86400(for 1 day) if `ServiceAffinity` == `ClientIP`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/service#timeout_seconds Service#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

