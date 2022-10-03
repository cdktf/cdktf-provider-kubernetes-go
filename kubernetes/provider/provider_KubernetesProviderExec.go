package provider


type KubernetesProviderExec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes#api_version KubernetesProvider#api_version}.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes#command KubernetesProvider#command}.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes#args KubernetesProvider#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes#env KubernetesProvider#env}.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

