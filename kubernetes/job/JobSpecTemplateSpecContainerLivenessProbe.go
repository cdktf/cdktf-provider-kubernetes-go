// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecTemplateSpecContainerLivenessProbe struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#exec Job#exec}
	Exec *JobSpecTemplateSpecContainerLivenessProbeExec `field:"optional" json:"exec" yaml:"exec"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#failure_threshold Job#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#grpc Job#grpc}
	Grpc interface{} `field:"optional" json:"grpc" yaml:"grpc"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#http_get Job#http_get}
	HttpGet *JobSpecTemplateSpecContainerLivenessProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#initial_delay_seconds Job#initial_delay_seconds}
	InitialDelaySeconds *float64 `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#period_seconds Job#period_seconds}
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#success_threshold Job#success_threshold}
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#tcp_socket Job#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
	// Number of seconds after which the probe times out. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/job#timeout_seconds Job#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

