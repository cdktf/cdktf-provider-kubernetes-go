// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type CronJobV1SpecJobTemplateSpecTemplateSpecVolumeAzureFile struct {
	// The name of secret that contains Azure Storage Account Name and Key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#secret_name CronJobV1#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Share Name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#share_name CronJobV1#share_name}
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#read_only CronJobV1#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The namespace of the secret that contains Azure Storage Account Name and Key.
	//
	// For Kubernetes up to 1.18.x the default is the same as the Pod. For Kubernetes 1.19.x and later the default is "default" namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#secret_namespace CronJobV1#secret_namespace}
	SecretNamespace *string `field:"optional" json:"secretNamespace" yaml:"secretNamespace"`
}
