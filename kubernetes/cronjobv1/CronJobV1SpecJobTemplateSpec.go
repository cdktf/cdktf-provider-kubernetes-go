// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1


type CronJobV1SpecJobTemplateSpec struct {
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#template CronJobV1#template}
	Template *CronJobV1SpecJobTemplateSpecTemplate `field:"required" json:"template" yaml:"template"`
	// Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers.
	//
	// Value must be a positive integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#active_deadline_seconds CronJobV1#active_deadline_seconds}
	ActiveDeadlineSeconds *float64 `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	// Specifies the number of retries before marking this job failed. Defaults to 6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#backoff_limit CronJobV1#backoff_limit}
	BackoffLimit *float64 `field:"optional" json:"backoffLimit" yaml:"backoffLimit"`
	// Specifies the limit for the number of retries within an index before marking this index as failed.
	//
	// When enabled the number of failures per index is kept in the pod's batch.kubernetes.io/job-index-failure-count annotation. It can only be set when Job's completionMode=Indexed, and the Pod's restart policy is Never. The field is immutable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#backoff_limit_per_index CronJobV1#backoff_limit_per_index}
	BackoffLimitPerIndex *float64 `field:"optional" json:"backoffLimitPerIndex" yaml:"backoffLimitPerIndex"`
	// Specifies how Pod completions are tracked. It can be `NonIndexed` (default) or `Indexed`. More info: https://kubernetes.io/docs/concepts/workloads/controllers/job/#completion-mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#completion_mode CronJobV1#completion_mode}
	CompletionMode *string `field:"optional" json:"completionMode" yaml:"completionMode"`
	// Specifies the desired number of successfully finished pods the job should be run with.
	//
	// Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value. Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#completions CronJobV1#completions}
	Completions *float64 `field:"optional" json:"completions" yaml:"completions"`
	// Controls generation of pod labels and pod selectors.
	//
	// Leave unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template. When true, the user is responsible for picking unique labels and specifying the selector. Failure to pick a unique label may cause this and other jobs to not function correctly. More info: https://git.k8s.io/community/contributors/design-proposals/selector-generation.md
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#manual_selector CronJobV1#manual_selector}
	ManualSelector interface{} `field:"optional" json:"manualSelector" yaml:"manualSelector"`
	// Controls generation of pod labels and pod selectors.
	//
	// Leave unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template. When true, the user is responsible for picking unique labels and specifying the selector. Failure to pick a unique label may cause this and other jobs to not function correctly. More info: https://git.k8s.io/community/contributors/design-proposals/selector-generation.md
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#max_failed_indexes CronJobV1#max_failed_indexes}
	MaxFailedIndexes *float64 `field:"optional" json:"maxFailedIndexes" yaml:"maxFailedIndexes"`
	// Specifies the maximum desired number of pods the job should run at any given time.
	//
	// The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#parallelism CronJobV1#parallelism}
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// pod_failure_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#pod_failure_policy CronJobV1#pod_failure_policy}
	PodFailurePolicy *CronJobV1SpecJobTemplateSpecPodFailurePolicy `field:"optional" json:"podFailurePolicy" yaml:"podFailurePolicy"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#selector CronJobV1#selector}
	Selector *CronJobV1SpecJobTemplateSpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed).
	//
	// If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/cron_job_v1#ttl_seconds_after_finished CronJobV1#ttl_seconds_after_finished}
	TtlSecondsAfterFinished *string `field:"optional" json:"ttlSecondsAfterFinished" yaml:"ttlSecondsAfterFinished"`
}

