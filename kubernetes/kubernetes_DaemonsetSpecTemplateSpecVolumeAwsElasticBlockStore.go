// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type DaemonsetSpecTemplateSpecVolumeAwsElasticBlockStore struct {
	// Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#volume_id Daemonset#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#fs_type Daemonset#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#partition Daemonset#partition}
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#read_only Daemonset#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}
