// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromResourceFieldRef struct {
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#resource ReplicationControllerV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#container_name ReplicationControllerV1#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#divisor ReplicationControllerV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}
