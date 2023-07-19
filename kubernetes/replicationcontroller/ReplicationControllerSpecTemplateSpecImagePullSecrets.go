package replicationcontroller


type ReplicationControllerSpecTemplateSpecImagePullSecrets struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/replication_controller#name ReplicationController#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

