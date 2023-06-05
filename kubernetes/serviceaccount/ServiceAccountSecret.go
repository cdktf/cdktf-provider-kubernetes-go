package serviceaccount


type ServiceAccountSecret struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/service_account#name ServiceAccount#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

