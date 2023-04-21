package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerm struct {
	// match_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job#match_expressions CronJob#match_expressions}
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
}

