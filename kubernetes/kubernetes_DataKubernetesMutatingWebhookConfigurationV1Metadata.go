// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type DataKubernetesMutatingWebhookConfigurationV1Metadata struct {
	// An unstructured key value map stored with the mutating webhook configuration that may be used to store arbitrary metadata.
	//
	// More info: http://kubernetes.io/docs/user-guide/annotations
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/mutating_webhook_configuration_v1#annotations DataKubernetesMutatingWebhookConfigurationV1#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the mutating webhook configuration.
	//
	// May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/mutating_webhook_configuration_v1#labels DataKubernetesMutatingWebhookConfigurationV1#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the mutating webhook configuration, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/d/mutating_webhook_configuration_v1#name DataKubernetesMutatingWebhookConfigurationV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

