package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeGrpc struct {
	// Number of the port to access on the container. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#port CronJobV1#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/cron_job_v1#service CronJobV1#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

