// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1Spec struct {
	// job_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#job_template CronJobV1#job_template}
	JobTemplate *CronJobV1SpecJobTemplate `field:"required" json:"jobTemplate" yaml:"jobTemplate"`
	// Cron format string, e.g. 0 * * * * or.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Specifies how to treat concurrent executions of a Job. Defaults to Allow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#concurrency_policy CronJobV1#concurrency_policy}
	ConcurrencyPolicy *string `field:"optional" json:"concurrencyPolicy" yaml:"concurrencyPolicy"`
	// The number of failed finished jobs to retain.
	//
	// This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#failed_jobs_history_limit CronJobV1#failed_jobs_history_limit}
	FailedJobsHistoryLimit *float64 `field:"optional" json:"failedJobsHistoryLimit" yaml:"failedJobsHistoryLimit"`
	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.
	//
	// Missed jobs executions will be counted as failed ones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#starting_deadline_seconds CronJobV1#starting_deadline_seconds}
	StartingDeadlineSeconds *float64 `field:"optional" json:"startingDeadlineSeconds" yaml:"startingDeadlineSeconds"`
	// The number of successful finished jobs to retain. Defaults to 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#successful_jobs_history_limit CronJobV1#successful_jobs_history_limit}
	SuccessfulJobsHistoryLimit *float64 `field:"optional" json:"successfulJobsHistoryLimit" yaml:"successfulJobsHistoryLimit"`
	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.
	//
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#suspend CronJobV1#suspend}
	Suspend interface{} `field:"optional" json:"suspend" yaml:"suspend"`
	// The time zone for the given schedule.
	//
	// If not specified, this will rely on the time zone of the kube-controller-manager process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#timezone CronJobV1#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

