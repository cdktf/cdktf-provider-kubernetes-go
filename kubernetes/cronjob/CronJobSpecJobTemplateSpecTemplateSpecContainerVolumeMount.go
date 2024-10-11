// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecContainerVolumeMount struct {
	// Path within the container at which the volume should be mounted. Must not contain ':'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#mount_path CronJob#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// This must match the Name of a Volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#name CronJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Mount propagation mode.
	//
	// mount_propagation determines how mounts are propagated from the host to container and the other way around. Valid values are None (default), HostToContainer and Bidirectional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#mount_propagation CronJob#mount_propagation}
	MountPropagation *string `field:"optional" json:"mountPropagation" yaml:"mountPropagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#read_only CronJob#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/cron_job#sub_path CronJob#sub_path}
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

