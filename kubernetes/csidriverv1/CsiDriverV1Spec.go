// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csidriverv1


type CsiDriverV1Spec struct {
	// Indicates if the CSI volume driver requires an attach operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/csi_driver_v1#attach_required CsiDriverV1#attach_required}
	AttachRequired interface{} `field:"required" json:"attachRequired" yaml:"attachRequired"`
	// Indicates that the CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/csi_driver_v1#pod_info_on_mount CsiDriverV1#pod_info_on_mount}
	PodInfoOnMount interface{} `field:"optional" json:"podInfoOnMount" yaml:"podInfoOnMount"`
	// Defines what kind of volumes this CSI volume driver supports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/csi_driver_v1#volume_lifecycle_modes CsiDriverV1#volume_lifecycle_modes}
	VolumeLifecycleModes *[]*string `field:"optional" json:"volumeLifecycleModes" yaml:"volumeLifecycleModes"`
}

