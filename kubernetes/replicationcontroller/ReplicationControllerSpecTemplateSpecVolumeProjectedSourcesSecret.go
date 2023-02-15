package replicationcontroller


type ReplicationControllerSpecTemplateSpecVolumeProjectedSourcesSecret struct {
	// items block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#items ReplicationController#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Name of the secret in the pod's namespace to use. More info: http://kubernetes.io/docs/user-guide/volumes#secrets.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#name ReplicationController#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional: Specify whether the Secret or it's keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#optional ReplicationController#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

