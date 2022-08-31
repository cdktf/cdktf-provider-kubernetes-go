// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PersistentVolumeClaimSpecResources struct {
	// Map describing the maximum amount of compute resources allowed. More info: http://kubernetes.io/docs/user-guide/compute-resources/.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_claim#limits PersistentVolumeClaim#limits}
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Map describing the minimum amount of compute resources required.
	//
	// If this is omitted for a container, it defaults to `limits` if that is explicitly specified, otherwise to an implementation-defined value. More info: http://kubernetes.io/docs/user-guide/compute-resources/
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_claim#requests PersistentVolumeClaim#requests}
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

