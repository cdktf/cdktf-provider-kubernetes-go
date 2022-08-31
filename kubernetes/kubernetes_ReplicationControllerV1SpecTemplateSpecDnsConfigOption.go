// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ReplicationControllerV1SpecTemplateSpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#name ReplicationControllerV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#value ReplicationControllerV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

