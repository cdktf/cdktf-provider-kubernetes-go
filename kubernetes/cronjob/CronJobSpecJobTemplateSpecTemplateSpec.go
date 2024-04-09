// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjob


type CronJobSpecJobTemplateSpecTemplateSpec struct {
	// Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers.
	//
	// Value must be a positive integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#active_deadline_seconds CronJob#active_deadline_seconds}
	ActiveDeadlineSeconds *float64 `field:"optional" json:"activeDeadlineSeconds" yaml:"activeDeadlineSeconds"`
	// affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#affinity CronJob#affinity}
	Affinity *CronJobSpecJobTemplateSpecTemplateSpecAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	// AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#automount_service_account_token CronJob#automount_service_account_token}
	AutomountServiceAccountToken interface{} `field:"optional" json:"automountServiceAccountToken" yaml:"automountServiceAccountToken"`
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#container CronJob#container}
	Container interface{} `field:"optional" json:"container" yaml:"container"`
	// dns_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#dns_config CronJob#dns_config}
	DnsConfig *CronJobSpecJobTemplateSpecTemplateSpecDnsConfig `field:"optional" json:"dnsConfig" yaml:"dnsConfig"`
	// Set DNS policy for containers within the pod.
	//
	// Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. Defaults to 'ClusterFirst'. More info: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#dns_policy CronJob#dns_policy}
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	// Enables generating environment variables for service discovery. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#enable_service_links CronJob#enable_service_links}
	EnableServiceLinks interface{} `field:"optional" json:"enableServiceLinks" yaml:"enableServiceLinks"`
	// host_aliases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#host_aliases CronJob#host_aliases}
	HostAliases interface{} `field:"optional" json:"hostAliases" yaml:"hostAliases"`
	// Use the host's ipc namespace. Optional: Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#host_ipc CronJob#host_ipc}
	HostIpc interface{} `field:"optional" json:"hostIpc" yaml:"hostIpc"`
	// Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#hostname CronJob#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Host networking requested for this pod.
	//
	// Use the host's network namespace. If this option is set, the ports that will be used must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#host_network CronJob#host_network}
	HostNetwork interface{} `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// Use the host's pid namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#host_pid CronJob#host_pid}
	HostPid interface{} `field:"optional" json:"hostPid" yaml:"hostPid"`
	// image_pull_secrets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#image_pull_secrets CronJob#image_pull_secrets}
	ImagePullSecrets interface{} `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	// init_container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#init_container CronJob#init_container}
	InitContainer interface{} `field:"optional" json:"initContainer" yaml:"initContainer"`
	// NodeName is a request to schedule this pod onto a specific node.
	//
	// If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#node_name CronJob#node_name}
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	//
	// Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#node_selector CronJob#node_selector}
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// os block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#os CronJob#os}
	Os *CronJobSpecJobTemplateSpecTemplateSpecOs `field:"optional" json:"os" yaml:"os"`
	// If specified, indicates the pod's priority.
	//
	// "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#priority_class_name CronJob#priority_class_name}
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// readiness_gate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#readiness_gate CronJob#readiness_gate}
	ReadinessGate interface{} `field:"optional" json:"readinessGate" yaml:"readinessGate"`
	// Restart policy for all containers within the pod. One of Always, OnFailure, Never. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#restart_policy CronJob#restart_policy}
	RestartPolicy *string `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// RuntimeClassName is a feature for selecting the container runtime configuration.
	//
	// The container runtime configuration is used to run a Pod's containers. More info: https://kubernetes.io/docs/concepts/containers/runtime-class
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#runtime_class_name CronJob#runtime_class_name}
	RuntimeClassName *string `field:"optional" json:"runtimeClassName" yaml:"runtimeClassName"`
	// If specified, the pod will be dispatched by specified scheduler.
	//
	// If not specified, the pod will be dispatched by default scheduler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#scheduler_name CronJob#scheduler_name}
	SchedulerName *string `field:"optional" json:"schedulerName" yaml:"schedulerName"`
	// security_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#security_context CronJob#security_context}
	SecurityContext *CronJobSpecJobTemplateSpecTemplateSpecSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	// ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: http://releases.k8s.io/HEAD/docs/design/service_accounts.md.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#service_account_name CronJob#service_account_name}
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// Share a single process namespace between all of the containers in a pod.
	//
	// When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#share_process_namespace CronJob#share_process_namespace}
	ShareProcessNamespace interface{} `field:"optional" json:"shareProcessNamespace" yaml:"shareProcessNamespace"`
	// If specified, the fully qualified Pod hostname will be "...svc.". If not specified, the pod will not have a domainname at all..
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#subdomain CronJob#subdomain}
	Subdomain *string `field:"optional" json:"subdomain" yaml:"subdomain"`
	// Optional duration in seconds the pod needs to terminate gracefully.
	//
	// May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#termination_grace_period_seconds CronJob#termination_grace_period_seconds}
	TerminationGracePeriodSeconds *float64 `field:"optional" json:"terminationGracePeriodSeconds" yaml:"terminationGracePeriodSeconds"`
	// toleration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#toleration CronJob#toleration}
	Toleration interface{} `field:"optional" json:"toleration" yaml:"toleration"`
	// topology_spread_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#topology_spread_constraint CronJob#topology_spread_constraint}
	TopologySpreadConstraint interface{} `field:"optional" json:"topologySpreadConstraint" yaml:"topologySpreadConstraint"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/cron_job#volume CronJob#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}

