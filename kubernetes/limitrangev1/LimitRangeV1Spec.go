package limitrangev1


type LimitRangeV1Spec struct {
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/limit_range_v1#limit LimitRangeV1#limit}
	Limit interface{} `field:"optional" json:"limit" yaml:"limit"`
}

