package k8s

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

)
func RealK8sClient(config *rest.Config, options client.Options) (client.Client, error) {
	return client.New(config, options)
}