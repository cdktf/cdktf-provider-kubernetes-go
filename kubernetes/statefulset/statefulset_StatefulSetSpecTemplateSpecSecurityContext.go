package statefulset


type StatefulSetSpecTemplateSpecSecurityContext struct {
	// A special supplemental group that applies to all containers in a pod.
	//
	// Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#fs_group StatefulSet#fs_group}
	FsGroup *string `field:"optional" json:"fsGroup" yaml:"fsGroup"`
	// The GID to run the entrypoint of the container process.
	//
	// Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#run_as_group StatefulSet#run_as_group}
	RunAsGroup *string `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#run_as_non_root StatefulSet#run_as_non_root}
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// The UID to run the entrypoint of the container process.
	//
	// Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#run_as_user StatefulSet#run_as_user}
	RunAsUser *string `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// seccomp_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#seccomp_profile StatefulSet#seccomp_profile}
	SeccompProfile *StatefulSetSpecTemplateSpecSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	// se_linux_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#se_linux_options StatefulSet#se_linux_options}
	SeLinuxOptions *StatefulSetSpecTemplateSpecSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
	// A list of groups applied to the first process run in each container, in addition to the container's primary GID.
	//
	// If unspecified, no groups will be added to any container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#supplemental_groups StatefulSet#supplemental_groups}
	SupplementalGroups *[]*float64 `field:"optional" json:"supplementalGroups" yaml:"supplementalGroups"`
	// sysctl block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#sysctl StatefulSet#sysctl}
	Sysctl interface{} `field:"optional" json:"sysctl" yaml:"sysctl"`
}
