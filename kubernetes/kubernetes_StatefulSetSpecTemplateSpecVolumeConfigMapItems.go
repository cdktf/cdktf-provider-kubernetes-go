// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type StatefulSetSpecTemplateSpecVolumeConfigMapItems struct {
	// The key to project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#key StatefulSet#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#mode StatefulSet#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The relative path of the file to map the key to.
	//
	// May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#path StatefulSet#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}
