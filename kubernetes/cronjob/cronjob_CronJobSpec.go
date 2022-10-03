package cronjob


type CronJobSpec struct {
	// job_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#job_template CronJob#job_template}
	JobTemplate *CronJobSpecJobTemplate `field:"required" json:"jobTemplate" yaml:"jobTemplate"`
	// Cron format string, e.g. 0 * * * * or @hourly, as schedule time of its jobs to be created and executed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#schedule CronJob#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// Specifies how to treat concurrent executions of a Job. Defaults to Allow.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#concurrency_policy CronJob#concurrency_policy}
	ConcurrencyPolicy *string `field:"optional" json:"concurrencyPolicy" yaml:"concurrencyPolicy"`
	// The number of failed finished jobs to retain.
	//
	// This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#failed_jobs_history_limit CronJob#failed_jobs_history_limit}
	FailedJobsHistoryLimit *float64 `field:"optional" json:"failedJobsHistoryLimit" yaml:"failedJobsHistoryLimit"`
	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.
	//
	// Missed jobs executions will be counted as failed ones.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#starting_deadline_seconds CronJob#starting_deadline_seconds}
	StartingDeadlineSeconds *float64 `field:"optional" json:"startingDeadlineSeconds" yaml:"startingDeadlineSeconds"`
	// The number of successful finished jobs to retain. Defaults to 3.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#successful_jobs_history_limit CronJob#successful_jobs_history_limit}
	SuccessfulJobsHistoryLimit *float64 `field:"optional" json:"successfulJobsHistoryLimit" yaml:"successfulJobsHistoryLimit"`
	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.
	//
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#suspend CronJob#suspend}
	Suspend interface{} `field:"optional" json:"suspend" yaml:"suspend"`
}

