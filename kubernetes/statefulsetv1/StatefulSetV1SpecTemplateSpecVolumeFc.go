// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecTemplateSpecVolumeFc struct {
	// FC target lun number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set_v1#lun StatefulSetV1#lun}
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// FC target worldwide names (WWNs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set_v1#target_ww_ns StatefulSetV1#target_ww_ns}
	TargetWwNs *[]*string `field:"required" json:"targetWwNs" yaml:"targetWwNs"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set_v1#fs_type StatefulSetV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set_v1#read_only StatefulSetV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

