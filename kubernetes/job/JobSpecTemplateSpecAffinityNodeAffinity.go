package job


type JobSpecTemplateSpecAffinityNodeAffinity struct {
	// preferred_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#preferred_during_scheduling_ignored_during_execution Job#preferred_during_scheduling_ignored_during_execution}
	PreferredDuringSchedulingIgnoredDuringExecution interface{} `field:"optional" json:"preferredDuringSchedulingIgnoredDuringExecution" yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
	// required_during_scheduling_ignored_during_execution block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#required_during_scheduling_ignored_during_execution Job#required_during_scheduling_ignored_during_execution}
	RequiredDuringSchedulingIgnoredDuringExecution *JobSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `field:"optional" json:"requiredDuringSchedulingIgnoredDuringExecution" yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
}

