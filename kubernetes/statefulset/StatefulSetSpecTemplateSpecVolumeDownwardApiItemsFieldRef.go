package statefulset


type StatefulSetSpecTemplateSpecVolumeDownwardApiItemsFieldRef struct {
	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#api_version StatefulSet#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Path of the field to select in the specified API version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#field_path StatefulSet#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
}

