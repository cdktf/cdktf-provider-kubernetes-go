package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecContainerSecurityContextSeLinuxOptions struct {
	// Level is SELinux level label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#level CronJobV1#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Role is a SELinux role label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#role CronJobV1#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Type is a SELinux type label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#type CronJobV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User is a SELinux user label that applies to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job_v1#user CronJobV1#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

