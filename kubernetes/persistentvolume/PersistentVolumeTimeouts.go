package persistentvolume


type PersistentVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/persistent_volume#create PersistentVolume#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}
