package resourcequota


type ResourceQuotaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#create ResourceQuota#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/resource_quota#update ResourceQuota#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

