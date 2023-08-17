package daemonset


type DaemonsetSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken struct {
	// Path specifies a relative path to the mount point of the projected volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemonset#path Daemonset#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Audience is the intended audience of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemonset#audience Daemonset#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// ExpirationSeconds is the expected duration of validity of the service account token.
	//
	// It defaults to 1 hour and must be at least 10 minutes (600 seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/daemonset#expiration_seconds Daemonset#expiration_seconds}
	ExpirationSeconds *float64 `field:"optional" json:"expirationSeconds" yaml:"expirationSeconds"`
}

