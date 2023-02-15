package persistentvolumeclaim


type PersistentVolumeClaimTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_claim#create PersistentVolumeClaim#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

