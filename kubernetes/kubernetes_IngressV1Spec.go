// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type IngressV1Spec struct {
	// default_backend block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#default_backend IngressV1#default_backend}
	DefaultBackend *IngressV1SpecDefaultBackend `field:"optional" json:"defaultBackend" yaml:"defaultBackend"`
	// IngressClassName is the name of the IngressClass cluster resource.
	//
	// The associated IngressClass defines which controller will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#ingress_class_name IngressV1#ingress_class_name}
	IngressClassName *string `field:"optional" json:"ingressClassName" yaml:"ingressClassName"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#rule IngressV1#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/ingress_v1#tls IngressV1#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

