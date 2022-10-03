//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (k *jsiiProxy_KubernetesProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateKubernetesProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_KubernetesProvider) validateSetExecParameters(val *KubernetesProviderExec) error {
	return nil
}

func (j *jsiiProxy_KubernetesProvider) validateSetExperimentsParameters(val *KubernetesProviderExperiments) error {
	return nil
}

func (j *jsiiProxy_KubernetesProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func validateNewKubernetesProviderParameters(scope constructs.Construct, id *string, config *KubernetesProviderConfig) error {
	return nil
}

