package podv1


type PodV1SpecVolumeNfs struct {
	// Path that is exported by the NFS server. More info: http://kubernetes.io/docs/user-guide/volumes#nfs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#path PodV1#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Server is the hostname or IP address of the NFS server. More info: http://kubernetes.io/docs/user-guide/volumes#nfs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#server PodV1#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. More info: http://kubernetes.io/docs/user-guide/volumes#nfs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#read_only PodV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

