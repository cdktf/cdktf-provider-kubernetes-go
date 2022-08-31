// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodV1SpecVolumeLocal struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_v1#path PodV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

