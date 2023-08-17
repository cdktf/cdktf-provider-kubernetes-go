package daemonset


type DaemonsetSpecTemplateSpecImagePullSecrets struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemonset#name Daemonset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

