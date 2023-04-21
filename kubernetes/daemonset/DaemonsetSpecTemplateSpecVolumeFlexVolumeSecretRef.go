package daemonset


type DaemonsetSpecTemplateSpecVolumeFlexVolumeSecretRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#name Daemonset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#namespace Daemonset#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

