package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecHostAliases struct {
	// Hostnames for the IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/cron_job_v1#hostnames CronJobV1#hostnames}
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// IP address of the host file entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/cron_job_v1#ip CronJobV1#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}
