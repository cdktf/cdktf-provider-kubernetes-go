// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerLifecyclePostStart struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cron_job_v1#exec CronJobV1#exec}
	Exec *CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cron_job_v1#http_get CronJobV1#http_get}
	HttpGet *CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/cron_job_v1#tcp_socket CronJobV1#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

