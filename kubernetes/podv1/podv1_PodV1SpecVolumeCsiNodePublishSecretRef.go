package podv1


type PodV1SpecVolumeCsiNodePublishSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#name PodV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

