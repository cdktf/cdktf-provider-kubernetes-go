package job


type JobSpecTemplateSpecVolumeFlocker struct {
	// Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job#dataset_name Job#dataset_name}
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// UUID of the dataset. This is unique identifier of a Flocker dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job#dataset_uuid Job#dataset_uuid}
	DatasetUuid *string `field:"optional" json:"datasetUuid" yaml:"datasetUuid"`
}
