package datakubernetesstorageclass


type DataKubernetesStorageClassAllowedTopologiesMatchLabelExpressions struct {
	// The label key that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/storage_class#key DataKubernetesStorageClass#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An array of string values. One value must match the label to be selected.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/storage_class#values DataKubernetesStorageClass#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

