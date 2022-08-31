// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type KubernetesProviderExperiments struct {
	// Enable the `kubernetes_manifest` resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes#manifest_resource KubernetesProvider#manifest_resource}
	ManifestResource interface{} `field:"optional" json:"manifestResource" yaml:"manifestResource"`
}

