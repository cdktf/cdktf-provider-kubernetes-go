package statefulsetv1


type StatefulSetV1SpecVolumeClaimTemplateSpec struct {
	// A set of the desired access modes the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#access_modes StatefulSetV1#access_modes}
	AccessModes *[]*string `field:"required" json:"accessModes" yaml:"accessModes"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#resources StatefulSetV1#resources}
	Resources *StatefulSetV1SpecVolumeClaimTemplateSpecResources `field:"required" json:"resources" yaml:"resources"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#selector StatefulSetV1#selector}
	Selector *StatefulSetV1SpecVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// Name of the storage class requested by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#storage_class_name StatefulSetV1#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// The binding reference to the PersistentVolume backing this claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set_v1#volume_name StatefulSetV1#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

