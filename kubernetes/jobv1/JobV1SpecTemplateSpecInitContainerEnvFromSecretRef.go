package jobv1


type JobV1SpecTemplateSpecInitContainerEnvFromSecretRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job_v1#name JobV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify whether the Secret must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/job_v1#optional JobV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}
