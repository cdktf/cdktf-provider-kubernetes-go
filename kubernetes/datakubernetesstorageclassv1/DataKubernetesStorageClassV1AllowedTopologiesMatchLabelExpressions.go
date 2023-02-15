package datakubernetesstorageclassv1


type DataKubernetesStorageClassV1AllowedTopologiesMatchLabelExpressions struct {
	// The label key that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/storage_class_v1#key DataKubernetesStorageClassV1#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An array of string values. One value must match the label to be selected.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/storage_class_v1#values DataKubernetesStorageClassV1#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

