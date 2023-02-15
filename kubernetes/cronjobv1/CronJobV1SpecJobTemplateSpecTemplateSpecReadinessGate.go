package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job_v1#condition_type CronJobV1#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

