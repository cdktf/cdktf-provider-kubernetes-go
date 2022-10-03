package poddisruptionbudget


type PodDisruptionBudgetSpec struct {
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_disruption_budget#selector PodDisruptionBudget#selector}
	Selector *PodDisruptionBudgetSpecSelector `field:"required" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_disruption_budget#max_unavailable PodDisruptionBudget#max_unavailable}.
	MaxUnavailable *string `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_disruption_budget#min_available PodDisruptionBudget#min_available}.
	MinAvailable *string `field:"optional" json:"minAvailable" yaml:"minAvailable"`
}

