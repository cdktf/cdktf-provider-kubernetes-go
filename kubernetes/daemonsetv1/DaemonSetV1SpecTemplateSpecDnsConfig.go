// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1


type DaemonSetV1SpecTemplateSpecDnsConfig struct {
	// A list of DNS name server IP addresses.
	//
	// This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemon_set_v1#nameservers DaemonSetV1#nameservers}
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	// option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemon_set_v1#option DaemonSetV1#option}
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// A list of DNS search domains for host-name lookup.
	//
	// This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemon_set_v1#searches DaemonSetV1#searches}
	Searches *[]*string `field:"optional" json:"searches" yaml:"searches"`
}

