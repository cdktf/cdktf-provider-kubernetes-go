package statefulset


type StatefulSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set#create StatefulSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set#delete StatefulSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set#read StatefulSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/stateful_set#update StatefulSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

