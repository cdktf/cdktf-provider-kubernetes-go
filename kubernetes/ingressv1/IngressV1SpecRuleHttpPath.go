package ingressv1


type IngressV1SpecRuleHttpPath struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/ingress_v1#backend IngressV1#backend}
	Backend *IngressV1SpecRuleHttpPathBackend `field:"optional" json:"backend" yaml:"backend"`
	// Path is matched against the path of an incoming request.
	//
	// Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value "Exact" or "Prefix".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/ingress_v1#path IngressV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// PathType determines the interpretation of the Path matching.
	//
	// PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is
	// done on a path element by element basis. A path element refers is the
	// list of labels in the path split by the '/' separator. A request is a
	// match for path p if every p is an element-wise prefix of p of the
	// request path. Note that if the last element of the path is a substring
	// of the last element in request path, it is not a match (e.g. /foo/bar
	// matches /foo/bar/baz, but does not match /foo/barbaz).
	// ImplementationSpecific: Interpretation of the Path matching is up to
	// the IngressClass. Implementations can treat this as a separate PathType
	// or treat it identically to Prefix or Exact path types.
	// Implementations are required to support all path types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/ingress_v1#path_type IngressV1#path_type}
	PathType *string `field:"optional" json:"pathType" yaml:"pathType"`
}

