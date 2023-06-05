package poddisruptionbudgetv1


type PodDisruptionBudgetV1Spec struct {
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/pod_disruption_budget_v1#selector PodDisruptionBudgetV1#selector}
	Selector *PodDisruptionBudgetV1SpecSelector `field:"required" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/pod_disruption_budget_v1#max_unavailable PodDisruptionBudgetV1#max_unavailable}.
	MaxUnavailable *string `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/pod_disruption_budget_v1#min_available PodDisruptionBudgetV1#min_available}.
	MinAvailable *string `field:"optional" json:"minAvailable" yaml:"minAvailable"`
}

