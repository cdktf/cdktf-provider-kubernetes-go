// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type DaemonsetSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution struct {
	// pod_affinity_term block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#pod_affinity_term Daemonset#pod_affinity_term}
	PodAffinityTerm *DaemonsetSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm `field:"required" json:"podAffinityTerm" yaml:"podAffinityTerm"`
	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#weight Daemonset#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}
