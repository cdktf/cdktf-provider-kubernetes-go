// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ReplicationControllerSpecTemplateSpecDnsConfig struct {
	// A list of DNS name server IP addresses.
	//
	// This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#nameservers ReplicationController#nameservers}
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	// option block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#option ReplicationController#option}
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// A list of DNS search domains for host-name lookup.
	//
	// This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#searches ReplicationController#searches}
	Searches *[]*string `field:"optional" json:"searches" yaml:"searches"`
}
