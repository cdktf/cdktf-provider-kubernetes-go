package configmapv1data


type ConfigMapV1DataMetadata struct {
	// The name of the ConfigMap.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/config_map_v1_data#name ConfigMapV1Data#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the ConfigMap.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/config_map_v1_data#namespace ConfigMapV1Data#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

