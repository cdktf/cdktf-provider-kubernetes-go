package datakubernetesstorageclass


type DataKubernetesStorageClassMetadata struct {
	// An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
	//
	// More info: http://kubernetes.io/docs/user-guide/annotations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/storage_class#annotations DataKubernetesStorageClass#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the storage class.
	//
	// May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/storage_class#labels DataKubernetesStorageClass#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the storage class, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/data-sources/storage_class#name DataKubernetesStorageClass#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

