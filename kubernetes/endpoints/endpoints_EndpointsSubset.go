package endpoints


type EndpointsSubset struct {
	// address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/endpoints#address Endpoints#address}
	Address interface{} `field:"optional" json:"address" yaml:"address"`
	// not_ready_address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/endpoints#not_ready_address Endpoints#not_ready_address}
	NotReadyAddress interface{} `field:"optional" json:"notReadyAddress" yaml:"notReadyAddress"`
	// port block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/endpoints#port Endpoints#port}
	Port interface{} `field:"optional" json:"port" yaml:"port"`
}

