package defaultserviceaccountv1


type DefaultServiceAccountV1ImagePullSecret struct {
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/default_service_account_v1#name DefaultServiceAccountV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

