package daemonset


type DaemonsetSpecTemplateSpecInitContainerEnvFromSecretRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#name Daemonset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the Secret must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#optional Daemonset#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

