package daemonset


type DaemonsetSpecTemplateSpecVolumeCephFs struct {
	// Monitors is a collection of Ceph monitors More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#monitors Daemonset#monitors}
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	// Used as the mounted root, rather than the full Ceph tree, default is /.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#path Daemonset#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to `false` (read/write). More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#read_only Daemonset#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The path to key ring for User, default is /etc/ceph/user.secret More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#secret_file Daemonset#secret_file}
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#secret_ref Daemonset#secret_ref}
	SecretRef *DaemonsetSpecTemplateSpecVolumeCephFsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// User is the rados user name, default is admin. More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#user Daemonset#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

