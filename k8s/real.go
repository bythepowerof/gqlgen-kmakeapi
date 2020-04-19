package k8s

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func RealK8sClient(config *rest.Config, options client.Options) (client.Client, error) {
	return client.New(config, options)
}

func RealK8sManager(config *rest.Config, options manager.Options) (manager.Manager, error) {
	return manager.New(config, options)
}
