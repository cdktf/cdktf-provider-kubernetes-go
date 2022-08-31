// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type JobV1SpecTemplateSpecInitContainerStartupProbeExec struct {
	// Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem.
	//
	// The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job_v1#command JobV1#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
}

