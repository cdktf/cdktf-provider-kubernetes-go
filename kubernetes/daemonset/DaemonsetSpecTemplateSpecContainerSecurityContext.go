// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecContainerSecurityContext struct {
	// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process.
	//
	// This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#allow_privilege_escalation Daemonset#allow_privilege_escalation}
	AllowPrivilegeEscalation interface{} `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#capabilities Daemonset#capabilities}
	Capabilities *DaemonsetSpecTemplateSpecContainerSecurityContextCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Run container in privileged mode.
	//
	// Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#privileged Daemonset#privileged}
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// Whether this container has a read-only root filesystem. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#read_only_root_filesystem Daemonset#read_only_root_filesystem}
	ReadOnlyRootFilesystem interface{} `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// The GID to run the entrypoint of the container process.
	//
	// Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#run_as_group Daemonset#run_as_group}
	RunAsGroup *string `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#run_as_non_root Daemonset#run_as_non_root}
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// The UID to run the entrypoint of the container process.
	//
	// Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#run_as_user Daemonset#run_as_user}
	RunAsUser *string `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// seccomp_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#seccomp_profile Daemonset#seccomp_profile}
	SeccompProfile *DaemonsetSpecTemplateSpecContainerSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	// se_linux_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/daemonset#se_linux_options Daemonset#se_linux_options}
	SeLinuxOptions *DaemonsetSpecTemplateSpecContainerSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
}

