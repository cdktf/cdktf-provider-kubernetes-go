// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type EndpointsV1SubsetAddress struct {
	// The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/endpoints_v1#ip EndpointsV1#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The Hostname of this endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/endpoints_v1#hostname EndpointsV1#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Node hosting this endpoint. This can be used to determine endpoints local to a node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/endpoints_v1#node_name EndpointsV1#node_name}
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
}
