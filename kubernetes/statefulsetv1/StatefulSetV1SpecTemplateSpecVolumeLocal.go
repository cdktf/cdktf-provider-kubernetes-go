package statefulsetv1


type StatefulSetV1SpecTemplateSpecVolumeLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/stateful_set_v1#path StatefulSetV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

