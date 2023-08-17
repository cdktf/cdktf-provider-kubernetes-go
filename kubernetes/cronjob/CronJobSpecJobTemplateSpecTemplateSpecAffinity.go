package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecAffinity struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job#node_affinity CronJob#node_affinity}
	NodeAffinity *CronJobSpecJobTemplateSpecTemplateSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// pod_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job#pod_affinity CronJob#pod_affinity}
	PodAffinity *CronJobSpecJobTemplateSpecTemplateSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_anti_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/cron_job#pod_anti_affinity CronJob#pod_anti_affinity}
	PodAntiAffinity *CronJobSpecJobTemplateSpecTemplateSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

