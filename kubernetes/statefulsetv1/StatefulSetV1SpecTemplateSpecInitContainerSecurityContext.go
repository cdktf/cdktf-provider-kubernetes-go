// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecTemplateSpecInitContainerSecurityContext struct {
	// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process.
	//
	// This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#allow_privilege_escalation StatefulSetV1#allow_privilege_escalation}
	AllowPrivilegeEscalation interface{} `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#capabilities StatefulSetV1#capabilities}
	Capabilities *StatefulSetV1SpecTemplateSpecInitContainerSecurityContextCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Run container in privileged mode.
	//
	// Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#privileged StatefulSetV1#privileged}
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// Whether this container has a read-only root filesystem. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#read_only_root_filesystem StatefulSetV1#read_only_root_filesystem}
	ReadOnlyRootFilesystem interface{} `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// The GID to run the entrypoint of the container process.
	//
	// Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#run_as_group StatefulSetV1#run_as_group}
	RunAsGroup *string `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#run_as_non_root StatefulSetV1#run_as_non_root}
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// The UID to run the entrypoint of the container process.
	//
	// Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#run_as_user StatefulSetV1#run_as_user}
	RunAsUser *string `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// seccomp_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#seccomp_profile StatefulSetV1#seccomp_profile}
	SeccompProfile *StatefulSetV1SpecTemplateSpecInitContainerSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	// se_linux_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/stateful_set_v1#se_linux_options StatefulSetV1#se_linux_options}
	SeLinuxOptions *StatefulSetV1SpecTemplateSpecInitContainerSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
}

