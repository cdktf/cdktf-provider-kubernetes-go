package provider


type KubernetesProviderExperiments struct {
	// Enable the `kubernetes_manifest` resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs#manifest_resource KubernetesProvider#manifest_resource}
	ManifestResource interface{} `field:"optional" json:"manifestResource" yaml:"manifestResource"`
}
