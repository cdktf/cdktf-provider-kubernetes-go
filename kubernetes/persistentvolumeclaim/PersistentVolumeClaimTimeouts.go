package persistentvolumeclaim


type PersistentVolumeClaimTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/persistent_volume_claim#create PersistentVolumeClaim#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

