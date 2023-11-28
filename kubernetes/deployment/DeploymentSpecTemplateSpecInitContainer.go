// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecInitContainer struct {
	// Name of the container specified as a DNS_LABEL.
	//
	// Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#name Deployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Arguments to the entrypoint.
	//
	// The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#args Deployment#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Entrypoint array.
	//
	// Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#command Deployment#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#env Deployment#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// env_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#env_from Deployment#env_from}
	EnvFrom interface{} `field:"optional" json:"envFrom" yaml:"envFrom"`
	// Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#image Deployment#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Image pull policy.
	//
	// One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images/#updating-images
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#image_pull_policy Deployment#image_pull_policy}
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// lifecycle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#lifecycle Deployment#lifecycle}
	Lifecycle *DeploymentSpecTemplateSpecInitContainerLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// liveness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#liveness_probe Deployment#liveness_probe}
	LivenessProbe *DeploymentSpecTemplateSpecInitContainerLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	// port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#port Deployment#port}
	Port interface{} `field:"optional" json:"port" yaml:"port"`
	// readiness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#readiness_probe Deployment#readiness_probe}
	ReadinessProbe *DeploymentSpecTemplateSpecInitContainerReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#resources Deployment#resources}
	Resources *DeploymentSpecTemplateSpecInitContainerResources `field:"optional" json:"resources" yaml:"resources"`
	// security_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#security_context Deployment#security_context}
	SecurityContext *DeploymentSpecTemplateSpecInitContainerSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	// startup_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#startup_probe Deployment#startup_probe}
	StartupProbe *DeploymentSpecTemplateSpecInitContainerStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	// Whether this container should allocate a buffer for stdin in the container runtime.
	//
	// If this is not set, reads from stdin in the container will always result in EOF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#stdin Deployment#stdin}
	Stdin interface{} `field:"optional" json:"stdin" yaml:"stdin"`
	// Whether the container runtime should close the stdin channel after it has been opened by a single attach.
	//
	// When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#stdin_once Deployment#stdin_once}
	StdinOnce interface{} `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	// Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.
	//
	// Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#termination_message_path Deployment#termination_message_path}
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	// Optional: Indicate how the termination message should be populated.
	//
	// File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#termination_message_policy Deployment#termination_message_policy}
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	// Whether this container should allocate a TTY for itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#tty Deployment#tty}
	Tty interface{} `field:"optional" json:"tty" yaml:"tty"`
	// volume_mount block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#volume_mount Deployment#volume_mount}
	VolumeMount interface{} `field:"optional" json:"volumeMount" yaml:"volumeMount"`
	// Container's working directory.
	//
	// If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#working_dir Deployment#working_dir}
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

