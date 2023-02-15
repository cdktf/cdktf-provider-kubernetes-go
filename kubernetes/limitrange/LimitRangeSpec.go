package limitrange


type LimitRangeSpec struct {
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/limit_range#limit LimitRange#limit}
	Limit interface{} `field:"optional" json:"limit" yaml:"limit"`
}

