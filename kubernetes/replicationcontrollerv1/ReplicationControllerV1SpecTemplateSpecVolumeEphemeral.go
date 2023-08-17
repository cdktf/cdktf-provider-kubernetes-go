package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecVolumeEphemeral struct {
	// volume_claim_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller_v1#volume_claim_template ReplicationControllerV1#volume_claim_template}
	VolumeClaimTemplate *ReplicationControllerV1SpecTemplateSpecVolumeEphemeralVolumeClaimTemplate `field:"required" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

