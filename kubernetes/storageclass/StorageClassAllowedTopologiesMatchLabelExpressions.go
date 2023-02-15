package storageclass


type StorageClassAllowedTopologiesMatchLabelExpressions struct {
	// The label key that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class#key StorageClass#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An array of string values. One value must match the label to be selected.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/storage_class#values StorageClass#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

