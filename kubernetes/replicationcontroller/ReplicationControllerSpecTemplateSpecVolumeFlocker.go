package replicationcontroller


type ReplicationControllerSpecTemplateSpecVolumeFlocker struct {
	// Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/replication_controller#dataset_name ReplicationController#dataset_name}
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// UUID of the dataset. This is unique identifier of a Flocker dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/replication_controller#dataset_uuid ReplicationController#dataset_uuid}
	DatasetUuid *string `field:"optional" json:"datasetUuid" yaml:"datasetUuid"`
}
