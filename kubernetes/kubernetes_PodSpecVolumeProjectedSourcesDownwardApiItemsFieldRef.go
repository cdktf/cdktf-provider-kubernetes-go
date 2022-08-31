// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSpecVolumeProjectedSourcesDownwardApiItemsFieldRef struct {
	// Version of the schema the FieldPath is written in terms of, defaults to 'v1'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#api_version Pod#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Path of the field to select in the specified API version.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#field_path Pod#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
}

