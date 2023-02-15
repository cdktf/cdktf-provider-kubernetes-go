package servicev1


type ServiceV1SpecSessionAffinityConfigClientIp struct {
	// Specifies the seconds of `ClientIP` type session sticky time.
	//
	// The value must be > 0 and <= 86400(for 1 day) if `ServiceAffinity` == `ClientIP`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/service_v1#timeout_seconds ServiceV1#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

