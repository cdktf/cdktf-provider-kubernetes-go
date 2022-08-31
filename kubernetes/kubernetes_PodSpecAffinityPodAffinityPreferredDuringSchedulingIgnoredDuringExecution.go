// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// pod_affinity_term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#pod_affinity_term Pod#pod_affinity_term}
	PodAffinityTerm *PodSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#weight Pod#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

