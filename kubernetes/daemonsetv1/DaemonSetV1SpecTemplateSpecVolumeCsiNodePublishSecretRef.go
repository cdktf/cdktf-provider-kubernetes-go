package daemonsetv1


type DaemonSetV1SpecTemplateSpecVolumeCsiNodePublishSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#name DaemonSetV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

