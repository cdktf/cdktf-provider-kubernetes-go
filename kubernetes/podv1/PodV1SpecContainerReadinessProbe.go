package podv1


type PodV1SpecContainerReadinessProbe struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#exec PodV1#exec}
	Exec *PodV1SpecContainerReadinessProbeExec `field:"optional" json:"exec" yaml:"exec"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#failure_threshold PodV1#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#grpc PodV1#grpc}
	Grpc interface{} `field:"optional" json:"grpc" yaml:"grpc"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#http_get PodV1#http_get}
	HttpGet *PodV1SpecContainerReadinessProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#initial_delay_seconds PodV1#initial_delay_seconds}
	InitialDelaySeconds *float64 `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#period_seconds PodV1#period_seconds}
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#success_threshold PodV1#success_threshold}
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#tcp_socket PodV1#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
	// Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_v1#timeout_seconds PodV1#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

