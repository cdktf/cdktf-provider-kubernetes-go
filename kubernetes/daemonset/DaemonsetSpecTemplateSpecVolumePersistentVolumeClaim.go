package daemonset


type DaemonsetSpecTemplateSpecVolumePersistentVolumeClaim struct {
	// ClaimName is the name of a PersistentVolumeClaim in the same.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemonset#claim_name Daemonset#claim_name}
	ClaimName *string `field:"optional" json:"claimName" yaml:"claimName"`
	// Will force the ReadOnly setting in VolumeMounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemonset#read_only Daemonset#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

