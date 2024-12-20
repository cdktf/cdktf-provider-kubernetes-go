// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podsecuritypolicyv1beta1


type PodSecurityPolicyV1Beta1Spec struct {
	// fs_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#fs_group PodSecurityPolicyV1Beta1#fs_group}
	FsGroup *PodSecurityPolicyV1Beta1SpecFsGroup `field:"required" json:"fsGroup" yaml:"fsGroup"`
	// run_as_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#run_as_user PodSecurityPolicyV1Beta1#run_as_user}
	RunAsUser *PodSecurityPolicyV1Beta1SpecRunAsUser `field:"required" json:"runAsUser" yaml:"runAsUser"`
	// supplemental_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#supplemental_groups PodSecurityPolicyV1Beta1#supplemental_groups}
	SupplementalGroups *PodSecurityPolicyV1Beta1SpecSupplementalGroups `field:"required" json:"supplementalGroups" yaml:"supplementalGroups"`
	// allowedCapabilities is a list of capabilities that can be requested to add to the container.
	//
	// Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#allowed_capabilities PodSecurityPolicyV1Beta1#allowed_capabilities}
	AllowedCapabilities *[]*string `field:"optional" json:"allowedCapabilities" yaml:"allowedCapabilities"`
	// allowed_flex_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#allowed_flex_volumes PodSecurityPolicyV1Beta1#allowed_flex_volumes}
	AllowedFlexVolumes interface{} `field:"optional" json:"allowedFlexVolumes" yaml:"allowedFlexVolumes"`
	// allowed_host_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#allowed_host_paths PodSecurityPolicyV1Beta1#allowed_host_paths}
	AllowedHostPaths interface{} `field:"optional" json:"allowedHostPaths" yaml:"allowedHostPaths"`
	// AllowedProcMountTypes is an allowlist of allowed ProcMountTypes.
	//
	// Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#allowed_proc_mount_types PodSecurityPolicyV1Beta1#allowed_proc_mount_types}
	AllowedProcMountTypes *[]*string `field:"optional" json:"allowedProcMountTypes" yaml:"allowedProcMountTypes"`
	// allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none.
	//
	// Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to allowlist all allowed unsafe sysctls explicitly to avoid rejection.
	//
	// Examples: e.g. "foo/*" allows "foo/bar", "foo/baz", etc. e.g. "foo.*" allows "foo.bar", "foo.baz", etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#allowed_unsafe_sysctls PodSecurityPolicyV1Beta1#allowed_unsafe_sysctls}
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#allow_privilege_escalation PodSecurityPolicyV1Beta1#allow_privilege_escalation}
	AllowPrivilegeEscalation interface{} `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.
	//
	// You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#default_add_capabilities PodSecurityPolicyV1Beta1#default_add_capabilities}
	DefaultAddCapabilities *[]*string `field:"optional" json:"defaultAddCapabilities" yaml:"defaultAddCapabilities"`
	// defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#default_allow_privilege_escalation PodSecurityPolicyV1Beta1#default_allow_privilege_escalation}
	DefaultAllowPrivilegeEscalation interface{} `field:"optional" json:"defaultAllowPrivilegeEscalation" yaml:"defaultAllowPrivilegeEscalation"`
	// forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none.
	//
	// Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.
	//
	// Examples: e.g. "foo/*" forbids "foo/bar", "foo/baz", etc. e.g. "foo.*" forbids "foo.bar", "foo.baz", etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#forbidden_sysctls PodSecurityPolicyV1Beta1#forbidden_sysctls}
	ForbiddenSysctls *[]*string `field:"optional" json:"forbiddenSysctls" yaml:"forbiddenSysctls"`
	// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#host_ipc PodSecurityPolicyV1Beta1#host_ipc}
	HostIpc interface{} `field:"optional" json:"hostIpc" yaml:"hostIpc"`
	// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#host_network PodSecurityPolicyV1Beta1#host_network}
	HostNetwork interface{} `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// hostPID determines if the policy allows the use of HostPID in the pod spec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#host_pid PodSecurityPolicyV1Beta1#host_pid}
	HostPid interface{} `field:"optional" json:"hostPid" yaml:"hostPid"`
	// host_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#host_ports PodSecurityPolicyV1Beta1#host_ports}
	HostPorts interface{} `field:"optional" json:"hostPorts" yaml:"hostPorts"`
	// privileged determines if a pod can request to be run as privileged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#privileged PodSecurityPolicyV1Beta1#privileged}
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.
	//
	// If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#read_only_root_filesystem PodSecurityPolicyV1Beta1#read_only_root_filesystem}
	ReadOnlyRootFilesystem interface{} `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// requiredDropCapabilities are the capabilities that will be dropped from the container.
	//
	// These are required to be dropped and cannot be added.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#required_drop_capabilities PodSecurityPolicyV1Beta1#required_drop_capabilities}
	RequiredDropCapabilities *[]*string `field:"optional" json:"requiredDropCapabilities" yaml:"requiredDropCapabilities"`
	// run_as_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#run_as_group PodSecurityPolicyV1Beta1#run_as_group}
	RunAsGroup *PodSecurityPolicyV1Beta1SpecRunAsGroup `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// se_linux block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#se_linux PodSecurityPolicyV1Beta1#se_linux}
	SeLinux *PodSecurityPolicyV1Beta1SpecSeLinux `field:"optional" json:"seLinux" yaml:"seLinux"`
	// volumes is an allowlist of volume plugins.
	//
	// Empty indicates that no volumes may be used. To allow all volumes you may use '*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_security_policy_v1beta1#volumes PodSecurityPolicyV1Beta1#volumes}
	Volumes *[]*string `field:"optional" json:"volumes" yaml:"volumes"`
}

