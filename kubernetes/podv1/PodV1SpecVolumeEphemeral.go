package podv1


type PodV1SpecVolumeEphemeral struct {
	// volume_claim_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod_v1#volume_claim_template PodV1#volume_claim_template}
	VolumeClaimTemplate *PodV1SpecVolumeEphemeralVolumeClaimTemplate `field:"required" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

