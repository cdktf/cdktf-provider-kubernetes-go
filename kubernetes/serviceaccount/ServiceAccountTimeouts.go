package serviceaccount


type ServiceAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/service_account#create ServiceAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

