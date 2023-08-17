package daemonset


type DaemonsetSpecTemplateSpecHostAliases struct {
	// Hostnames for the IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemonset#hostnames Daemonset#hostnames}
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// IP address of the host file entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemonset#ip Daemonset#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

