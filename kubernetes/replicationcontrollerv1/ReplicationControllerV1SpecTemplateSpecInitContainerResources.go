package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecInitContainerResources struct {
	// Describes the maximum amount of compute resources allowed. More info: http://kubernetes.io/docs/user-guide/compute-resources/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller_v1#limits ReplicationControllerV1#limits}
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/replication_controller_v1#requests ReplicationControllerV1#requests}
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

