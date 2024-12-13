// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolumeGlusterfs struct {
	// The endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job#endpoints_name CronJob#endpoints_name}
	EndpointsName *string `field:"required" json:"endpointsName" yaml:"endpointsName"`
	// The Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job#path CronJob#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/cron_job#read_only CronJob#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

