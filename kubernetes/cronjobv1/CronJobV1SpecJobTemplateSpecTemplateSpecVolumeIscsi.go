package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecVolumeIscsi struct {
	// Target iSCSI Qualified Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#iqn CronJobV1#iqn}
	Iqn *string `field:"required" json:"iqn" yaml:"iqn"`
	// iSCSI target portal.
	//
	// The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#target_portal CronJobV1#target_portal}
	TargetPortal *string `field:"required" json:"targetPortal" yaml:"targetPortal"`
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#iscsi
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#fs_type CronJobV1#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// iSCSI interface name that uses an iSCSI transport. Defaults to 'default' (tcp).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#iscsi_interface CronJobV1#iscsi_interface}
	IscsiInterface *string `field:"optional" json:"iscsiInterface" yaml:"iscsiInterface"`
	// iSCSI target lun number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#lun CronJobV1#lun}
	Lun *float64 `field:"optional" json:"lun" yaml:"lun"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#read_only CronJobV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

