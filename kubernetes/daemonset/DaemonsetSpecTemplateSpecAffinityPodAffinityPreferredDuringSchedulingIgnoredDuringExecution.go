package daemonset


type DaemonsetSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// pod_affinity_term block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemonset#pod_affinity_term Daemonset#pod_affinity_term}
	PodAffinityTerm *DaemonsetSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemonset#weight Daemonset#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}
