// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type DeploymentV1Spec struct {
	// template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#template DeploymentV1#template}
	Template *DeploymentV1SpecTemplate `field:"required" json:"template" yaml:"template"`
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available.
	//
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#min_ready_seconds DeploymentV1#min_ready_seconds}
	MinReadySeconds *float64 `field:"optional" json:"minReadySeconds" yaml:"minReadySeconds"`
	// Indicates that the deployment is paused.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#paused DeploymentV1#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// The maximum time in seconds for a deployment to make progress before it is considered to be failed.
	//
	// The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#progress_deadline_seconds DeploymentV1#progress_deadline_seconds}
	ProgressDeadlineSeconds *float64 `field:"optional" json:"progressDeadlineSeconds" yaml:"progressDeadlineSeconds"`
	// Number of desired pods. This is a string to be able to distinguish between explicit zero and not specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#replicas DeploymentV1#replicas}
	Replicas *string `field:"optional" json:"replicas" yaml:"replicas"`
	// The number of old ReplicaSets to retain to allow rollback.
	//
	// This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#revision_history_limit DeploymentV1#revision_history_limit}
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#selector DeploymentV1#selector}
	Selector *DeploymentV1SpecSelector `field:"optional" json:"selector" yaml:"selector"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#strategy DeploymentV1#strategy}
	Strategy *DeploymentV1SpecStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

