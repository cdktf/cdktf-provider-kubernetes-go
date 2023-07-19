package cronjob


type CronJobSpecJobTemplateSpec struct {
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#template CronJob#template}
	Template *CronJobSpecJobTemplateSpecTemplate `field:"required" json:"template" yaml:"template"`
	// Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers.
	//
	// Value must be a positive integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#active_deadline_seconds CronJob#active_deadline_seconds}
	ActiveDeadlineSeconds *float64 `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	// Specifies the number of retries before marking this job failed. Defaults to 6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#backoff_limit CronJob#backoff_limit}
	BackoffLimit *float64 `field:"optional" json:"backoffLimit" yaml:"backoffLimit"`
	// Specifies how Pod completions are tracked. It can be `NonIndexed` (default) or `Indexed`. For more information: https://kubernetes.io/docs/concepts/workloads/controllers/job/#completion-mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#completion_mode CronJob#completion_mode}
	CompletionMode *string `field:"optional" json:"completionMode" yaml:"completionMode"`
	// Specifies the desired number of successfully finished pods the job should be run with.
	//
	// Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value. Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#completions CronJob#completions}
	Completions *float64 `field:"optional" json:"completions" yaml:"completions"`
	// Controls generation of pod labels and pod selectors.
	//
	// Leave unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template. When true, the user is responsible for picking unique labels and specifying the selector. Failure to pick a unique label may cause this and other jobs to not function correctly. More info: https://git.k8s.io/community/contributors/design-proposals/selector-generation.md
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#manual_selector CronJob#manual_selector}
	ManualSelector interface{} `field:"optional" json:"manualSelector" yaml:"manualSelector"`
	// Specifies the maximum desired number of pods the job should run at any given time.
	//
	// The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#parallelism CronJob#parallelism}
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#selector CronJob#selector}
	Selector *CronJobSpecJobTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed).
	//
	// If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/cron_job#ttl_seconds_after_finished CronJob#ttl_seconds_after_finished}
	TtlSecondsAfterFinished *string `field:"optional" json:"ttlSecondsAfterFinished" yaml:"ttlSecondsAfterFinished"`
}

