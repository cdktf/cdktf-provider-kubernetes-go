package statefulset


type StatefulSetSpecTemplateSpecVolumeSecret struct {
	// Optional: mode bits to use on created files by default.
	//
	// Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#default_mode StatefulSet#default_mode}
	DefaultMode *string `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#items StatefulSet#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Optional: Specify whether the Secret or its keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#optional StatefulSet#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
	// Name of the secret in the pod's namespace to use. More info: http://kubernetes.io/docs/user-guide/volumes#secrets.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#secret_name StatefulSet#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

