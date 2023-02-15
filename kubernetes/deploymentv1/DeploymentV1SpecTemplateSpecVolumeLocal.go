package deploymentv1


type DeploymentV1SpecTemplateSpecVolumeLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#path DeploymentV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

