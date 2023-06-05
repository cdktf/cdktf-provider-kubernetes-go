package replicationcontroller


type ReplicationControllerSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef struct {
	// Version of the schema the FieldPath is written in terms of, defaults to 'v1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/replication_controller#api_version ReplicationController#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Path of the field to select in the specified API version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/replication_controller#field_path ReplicationController#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
}

