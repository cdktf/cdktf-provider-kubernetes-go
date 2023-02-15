package daemonsetv1


type DaemonSetV1SpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// pod_affinity_term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#pod_affinity_term DaemonSetV1#pod_affinity_term}
	PodAffinityTerm *DaemonSetV1SpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#weight DaemonSetV1#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

