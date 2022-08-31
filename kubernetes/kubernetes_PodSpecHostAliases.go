// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSpecHostAliases struct {
	// Hostnames for the IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#hostnames Pod#hostnames}
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// IP address of the host file entry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#ip Pod#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

