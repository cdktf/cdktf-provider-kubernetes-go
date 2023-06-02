package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecVolumeHostPath struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#hostpath.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/replication_controller_v1#path ReplicationControllerV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Type for HostPath volume. Allowed values are "" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/replication_controller_v1#type ReplicationControllerV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

