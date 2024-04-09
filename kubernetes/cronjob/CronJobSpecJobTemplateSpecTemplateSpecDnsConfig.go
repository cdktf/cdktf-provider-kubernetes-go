// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecDnsConfig struct {
	// A list of DNS name server IP addresses.
	//
	// This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#nameservers CronJob#nameservers}
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	// option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#option CronJob#option}
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// A list of DNS search domains for host-name lookup.
	//
	// This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#searches CronJob#searches}
	Searches *[]*string `field:"optional" json:"searches" yaml:"searches"`
}

