package daemonset


type DaemonsetSpecTemplateSpecVolumeHostPath struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#hostpath.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#path Daemonset#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Type for HostPath volume. Allowed values are "" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#type Daemonset#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

