// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodV1SpecContainerVolumeMount struct {
	// Path within the container at which the volume should be mounted. Must not contain ':'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#mount_path PodV1#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// This must match the Name of a Volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#name PodV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Mount propagation mode.
	//
	// mount_propagation determines how mounts are propagated from the host to container and the other way around. Valid values are None (default), HostToContainer and Bidirectional.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#mount_propagation PodV1#mount_propagation}
	MountPropagation *string `field:"optional" json:"mountPropagation" yaml:"mountPropagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#read_only PodV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#sub_path PodV1#sub_path}
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}
