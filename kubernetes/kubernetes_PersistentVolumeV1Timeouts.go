// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PersistentVolumeV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume_v1#create PersistentVolumeV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

