package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecContainerEnvFromConfigMapRef struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#name ReplicationControllerV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the ConfigMap must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#optional ReplicationControllerV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

