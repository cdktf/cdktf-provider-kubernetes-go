package pod


type PodSpecInitContainerStartupProbeHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#name Pod#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#value Pod#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

