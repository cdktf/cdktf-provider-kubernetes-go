package pod


type PodSpecImagePullSecrets struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/pod#name Pod#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

