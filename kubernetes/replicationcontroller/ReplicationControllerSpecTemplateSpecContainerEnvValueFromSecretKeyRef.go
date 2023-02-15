package replicationcontroller


type ReplicationControllerSpecTemplateSpecContainerEnvValueFromSecretKeyRef struct {
	// The key of the secret to select from. Must be a valid secret key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#key ReplicationController#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#name ReplicationController#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the Secret or its key must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#optional ReplicationController#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

