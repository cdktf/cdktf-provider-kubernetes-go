package limitrangev1


type LimitRangeV1Spec struct {
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/limit_range_v1#limit LimitRangeV1#limit}
	Limit interface{} `field:"optional" json:"limit" yaml:"limit"`
}

