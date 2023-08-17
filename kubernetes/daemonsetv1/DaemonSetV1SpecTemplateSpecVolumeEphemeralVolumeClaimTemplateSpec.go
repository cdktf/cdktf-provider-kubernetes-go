package daemonsetv1


type DaemonSetV1SpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec struct {
	// A set of the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemon_set_v1#access_modes DaemonSetV1#access_modes}
	AccessModes *[]*string `field:"required" json:"accessModes" yaml:"accessModes"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemon_set_v1#resources DaemonSetV1#resources}
	Resources *DaemonSetV1SpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResources `field:"required" json:"resources" yaml:"resources"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemon_set_v1#selector DaemonSetV1#selector}
	Selector *DaemonSetV1SpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// Name of the storage class requested by the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemon_set_v1#storage_class_name DaemonSetV1#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// The binding reference to the PersistentVolume backing this claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemon_set_v1#volume_name DaemonSetV1#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

